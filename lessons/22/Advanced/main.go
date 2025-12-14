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

// Змінна для керування маршрутизацією
var (
	tmplDir     = "templates/"
	store       *sessions.CookieStore
	db          *sql.DB
	csrfProtect func(http.Handler) http.Handler
)

const sessionName = "queue-session"

// Структура для клієнтів черги
type User struct {
	ID           int
	Username     string
	PasswordHash string
	CreatedAt    time.Time
}

// Функція для відображення сторінок
func renderTemplate(w http.ResponseWriter, r *http.Request, name string, data map[string]interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}

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

// Обробник маршруту /register для сторінки реєстрації
func registerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//Обробка запитів
	case http.MethodGet:
		renderTemplate(w, r, "register.html", nil)
	case http.MethodPost:
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" || password == "" {
			http.Error(w, "Fill all fields", http.StatusBadRequest)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Hash error", http.StatusInternalServerError)
			return
		}

		//Запис у MySQL
		_, err = db.Exec("INSERT INTO users (username, password_hash) VALUES (?, ?)", username, string(hash))
		if err != nil {
			http.Error(w, "Failed to register user: "+err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("Registering user:")
		fmt.Println("Username:", username)
		fmt.Println("Password hash:", string(hash))

		renderTemplate(w, r, "register_success.html", nil)
	default:
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
	}
}

// Обробник маршруту /lodin для сторінки входу
func loginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//Обробка запитів
	case http.MethodGet:
		renderTemplate(w, r, "login.html", nil)
	case http.MethodPost:
		username := r.FormValue("username")
		password := r.FormValue("password")
		if username == "" || password == "" {
			http.Error(w, "Заповніть всі поля", http.StatusBadRequest)
			return
		}

		var id int
		var passHash string
		//Запис у MySQL
		err := db.QueryRow("SELECT id, password_hash FROM users WHERE username = ?", username).Scan(&id, &passHash)
		if err == sql.ErrNoRows {
			http.Error(w, "Account data is corrupted", http.StatusUnauthorized)
			return
		} else if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(passHash), []byte(password))
		if err != nil {
			http.Error(w, "Incorrect account data", http.StatusUnauthorized)
			return
		}

		session, _ := store.Get(r, sessionName)
		session.Values["authenticated"] = true
		session.Values["username"] = username
		session.Save(r, w)

		http.Redirect(w, r, "/home", http.StatusSeeOther)
	default:
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
	}
}

// Обробник маршруту /register для головної сторінки (сторінки після успішної реєстрації)
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

// Обробник маршруту /logout
// На відмінну від інших маршрутів є перехідним процесом
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, sessionName)
	session.Values["authenticated"] = false
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func main() {
	//Перевірка встановлення DB_DSN
	//Для встановленн необхідно виконати в командному рядку:
	//set DB_DSN=postgres://myuser:mypassword@localhost:5432/mydatabase?sslmode=disable
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN not installed")
	}

	//Перевірка встановлення ключа сертифікату
	//Для встановленн необхідно виконати в командному рядку:
	//set CSRF_AUTH_KEY=your-32-character-csrf-key-here-minimum
	csrfKey := os.Getenv("CSRF_AUTH_KEY")
	if csrfKey == "" || len(csrfKey) < 32 {
		log.Fatal("CSRF_AUTH_KEY must be installed and have size at least 32 bites")
	}

	//Перевірка встановлення ключа сертифікату
	//Для встановленн необхідно виконати в командному рядку:
	//set SESSION_KEY=your-32-character-session-key-here-minimum
	sessionKey := os.Getenv("SESSION_KEY")
	if sessionKey == "" || len(sessionKey) < 32 {
		log.Fatal("SESSION_KEY must be installed and have size at least 32 bites")
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

	var err error
	//Підключення до бази даних
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Помилка підключення до БД: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("DB ping error: %v", err)
	}

	//Створення файлів Cookie, щоб зберігати дані входу користувача
	store = sessions.NewCookieStore([]byte(sessionKey))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	//Створення оболонки захисту
	csrfProtect = csrf.Protect([]byte(csrfKey), csrf.Secure(true))

	//Встановлення маршрутизації
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	})
	mux.HandleFunc("/register", registerHandler)
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/home", homeHandler)
	mux.HandleFunc("/logout", logoutHandler)

	//Запуск сервера
	handler := csrfProtect(mux)

	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Starting HTTPS server on https://localhost:%s\n", port)
	fmt.Println("Sertificate: ", certFile, " key: ", keyFile)
	if err := http.ListenAndServeTLS(addr, certFile, keyFile, handler); err != nil {
		log.Fatalf("ListenAndServeTLS: %v", err)
	}
}
