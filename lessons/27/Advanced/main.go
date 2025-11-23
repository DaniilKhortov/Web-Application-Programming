package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var (
	tmplDir     = "templates/"
	store       *sessions.CookieStore
	db          *sql.DB
	csrfProtect func(http.Handler) http.Handler
)

const sessionName = "queue-session"

// User model
type User struct {
	ID           int
	Username     string
	PasswordHash string
	CreatedAt    time.Time
}

// helper: render template with CSRF token
func renderTemplate(w http.ResponseWriter, r *http.Request, name string, data map[string]interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}
	// inject CSRF token into template data
	data["CSRF"] = csrf.TemplateField(r)
	t, err := template.ParseFiles(tmplDir+name, tmplDir+"base.html")
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err := t.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, "Template exec error: "+err.Error(), http.StatusInternalServerError)
	}
}

// Registration page (GET) and handler (POST)
func registerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		renderTemplate(w, r, "register.html", nil)
	case http.MethodPost:
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" || password == "" {
			http.Error(w, "Заповніть всі поля", http.StatusBadRequest)
			return
		}

		// hash password
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Помилка хешування", http.StatusInternalServerError)
			return
		}

		// insert into DB (use prepared statement to avoid injection)
		_, err = db.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", username, string(hash))
		if err != nil {
			// unique username violation -> user-friendly
			http.Error(w, "Не вдалося створити користувача: "+err.Error(), http.StatusBadRequest)
			return
		}

		// log to server console (username + hash)
		fmt.Println("=== New user registered ===")
		fmt.Println("Username:", username)
		fmt.Println("Password hash:", string(hash))
		fmt.Println("===========================")

		renderTemplate(w, r, "register_success.html", nil)
	default:
		http.Error(w, "Метод не підтримується", http.StatusMethodNotAllowed)
	}
}

// Login handler
func loginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		renderTemplate(w, r, "login.html", nil)
	case http.MethodPost:
		username := r.FormValue("username")
		password := r.FormValue("password")
		if username == "" || password == "" {
			http.Error(w, "Заповніть всі поля", http.StatusBadRequest)
			return
		}

		// fetch user hash from DB
		var id int
		var passHash string
		err := db.QueryRow("SELECT id, password_hash FROM users WHERE username = $1", username).Scan(&id, &passHash)
		if err == sql.ErrNoRows {
			http.Error(w, "Невірні облікові дані", http.StatusUnauthorized)
			return
		} else if err != nil {
			http.Error(w, "Помилка сервера", http.StatusInternalServerError)
			return
		}

		// compare hash
		err = bcrypt.CompareHashAndPassword([]byte(passHash), []byte(password))
		if err != nil {
			http.Error(w, "Невірні облікові дані", http.StatusUnauthorized)
			return
		}

		// create session
		session, _ := store.Get(r, sessionName)
		session.Values["authenticated"] = true
		session.Values["username"] = username
		session.Save(r, w)

		http.Redirect(w, r, "/home", http.StatusSeeOther)
	default:
		http.Error(w, "Метод не підтримується", http.StatusMethodNotAllowed)
	}
}

// Home page (requires auth)
func homeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, sessionName)
	auth, _ := session.Values["authenticated"].(bool)
	if !auth {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	username, _ := session.Values["username"].(string)
	data := map[string]interface{}{
		"Username": username,
	}
	renderTemplate(w, r, "home.html", data)
}

// Logout
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, sessionName)
	session.Values["authenticated"] = false
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func main() {
	// read config from env
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN не встановлено")
	}

	csrfKey := os.Getenv("CSRF_AUTH_KEY")
	if csrfKey == "" || len(csrfKey) < 32 {
		log.Fatal("CSRF_AUTH_KEY має бути встановлено й бути принаймні 32 байти")
	}

	sessionKey := os.Getenv("SESSION_KEY")
	if sessionKey == "" || len(sessionKey) < 32 {
		log.Fatal("SESSION_KEY має бути встановлено й бути принаймні 32 байти")
	}

	certFile := os.Getenv("CERT_FILE")
	if certFile == "" {
		certFile = "server.crt"
	}
	keyFile := os.Getenv("KEY_FILE")
	if keyFile == "" {
		keyFile = "server.key"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "4430"
	}

	// open DB
	var err error
	db, err = sql.Open("postgres", dsn) // якщо використовуєш MySQL — змінити драйвер та DSN
	if err != nil {
		log.Fatalf("Помилка підключення до БД: %v", err)
	}
	defer db.Close()

	// simple ping
	if err := db.Ping(); err != nil {
		log.Fatalf("DB ping error: %v", err)
	}

	// cookie store
	store = sessions.NewCookieStore([]byte(sessionKey))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 1, // 1 day
		HttpOnly: true,
		Secure:   true, // в локальному тесті https повинен бути увімкнений (ми використовуємо ListenAndServeTLS)
		SameSite: http.SameSiteStrictMode,
	}

	// CSRF protection middleware
	csrfProtect = csrf.Protect([]byte(csrfKey), csrf.Secure(true))

	// routes
	mux := http.NewServeMux()
	mux.HandleFunc("/register", registerHandler)
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/home", homeHandler)
	mux.HandleFunc("/logout", logoutHandler)

	// static (if потрібні стилі) — приклад
	// mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// In production add timeouts; for demo we keep simple
	handler := csrfProtect(mux)

	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Starting HTTPS server on https://localhost:%s\n", port)
	fmt.Println("Сертифікат: ", certFile, " ключ: ", keyFile)
	if err := http.ListenAndServeTLS(addr, certFile, keyFile, handler); err != nil {
		log.Fatalf("ListenAndServeTLS: %v", err)
	}
}
