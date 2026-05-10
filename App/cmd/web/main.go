package main

import (
	"log"
	"net/http"
	"os"

	"QueueApp/internal/data"
	"QueueApp/internal/db"
	handler "QueueApp/internal/handlers"
	"log/slog"

	"github.com/gorilla/csrf"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//Створення JSON логера
	Logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	Data := data.Lessons
	var DB *gorm.DB
	var err error

	slog.SetDefault(Logger)

	app := &handler.App{
		Logger: Logger,
		Data:   Data,
		DB:     DB,
	}

	app.DB, err = db.InitDB(Logger)
	if err != nil {
		log.Fatal(err)
	}

	//Створення маршрутизатора
	mux := http.NewServeMux()

	//Створення обробника static для файлу style.css, що знаходиться у /static
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	//Створення кореневого обробника
	mux.HandleFunc("GET /{$}", app.LandingHandler)

	//Створення  обробника головної сторінки
	mux.HandleFunc("GET /main", app.AuthMiddleware(app.HomeHandler))

	//Створення  обробника для завантаження файлу до репозиторію
	mux.HandleFunc("POST /upload", app.AuthMiddleware(app.UploadHandler))

	//Створення  обробника для отримання відомостей про оцінку
	mux.HandleFunc("GET /taskStatus", app.TaskStatusHandler)

	//Створення обробника для шляху GET /register
	mux.HandleFunc("GET /register", app.RegisterPageHandler)

	//Створення обробника для шляху POST /register
	mux.HandleFunc("POST /register", app.RegisterSubmitHandler)

	//Створення обробника для шляху GET /login
	mux.HandleFunc("GET /login", app.LoginPageHandler)

	//Створення обробника для шляху POST /login
	mux.HandleFunc("POST /login", app.LoginSubmitHandler)

	//Створення обробника для шляху GET /admin
	mux.HandleFunc("GET /admin", app.AdminMiddleware(app.AdminPageHandler))

	//Створення обробника для шляху PUT  /admin/update-mark
	mux.HandleFunc("PUT /admin/update-mark/{id}", app.AdminMiddleware(app.UpdateMarkHandler))

	//Створення обробника для шляху PUT  /admin/update-role
	mux.HandleFunc("PUT /admin/update-role/{id}", app.AdminMiddleware(app.UpdateRoleHandler))

	//Створення обробника для шляху GET /lesson1
	mux.HandleFunc("GET /lesson1", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 1, "lesson1") }))
	//Створення обробника для шляху GET /lesson2
	mux.HandleFunc("GET /lesson2", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 2, "lesson2") }))

	//Створення обробника для шляху GET /lesson3
	mux.HandleFunc("GET /lesson3", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 3, "lesson3") }))

	//Створення обробника для шляху GET /lesson4
	mux.HandleFunc("GET /lesson4", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 4, "lesson4") }))

	//Створення обробника для шляху GET /lesson5
	mux.HandleFunc("GET /lesson5", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 5, "lesson5") }))

	//Створення обробника для шляху GET /lesson6
	mux.HandleFunc("GET /lesson6", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 6, "lesson6") }))

	//Створення обробника для шляху GET /lesson7
	mux.HandleFunc("GET /lesson7", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 7, "lesson7") }))

	//Створення обробника для шляху GET /lesson8
	mux.HandleFunc("GET /lesson8", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 8, "lesson8") }))

	//Створення обробника для шляху GET /lesson9
	mux.HandleFunc("GET /lesson9", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 9, "lesson9") }))

	//Створення обробника для шляху GET /lesson10
	mux.HandleFunc("GET /lesson10", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 10, "lesson10") }))

	//Створення обробника для шляху GET /lesson11
	mux.HandleFunc("GET /lesson11", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 11, "lesson11") }))

	//Створення обробника для шляху GET /lesson12
	mux.HandleFunc("GET /lesson12", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 12, "lesson12") }))

	//Створення обробника для шляху GET /lesson13
	mux.HandleFunc("GET /lesson13", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 13, "lesson13") }))
	//Створення обробника для шляху GET /lesson14
	mux.HandleFunc("GET /lesson14", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 14, "lesson14") }))

	//Створення обробника для шляху GET /lesson15
	mux.HandleFunc("GET /lesson15", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 15, "lesson15") }))
	//Створення обробника для шляху GET /lesson16
	mux.HandleFunc("GET /lesson16", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 16, "lesson16") }))
	//Створення обробника для шляху app.AuthMiddleware(GET /lesson17
	mux.HandleFunc("GET /lesson17", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 17, "lesson17") }))
	//Створення обробника для шляху app.AuthMiddleware(GET /lesson18
	mux.HandleFunc("GET /lesson18", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 18, "lesson18") }))
	//Створення обробника для шляху GET /lesson19
	mux.HandleFunc("GET /lesson19", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 19, "lesson19") }))
	//Створення обробника для шляху GET /lesson20
	mux.HandleFunc("GET /lesson20", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 20, "lesson20") }))
	//Створення обробника для шляху GET /lesson21
	mux.HandleFunc("GET /lesson21", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 21, "lesson21") }))
	//Створення обробника для шляху GET /lesson22
	mux.HandleFunc("GET /lesson22", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 22, "lesson22") }))
	//Створення обробника для шляху GET /lesson23
	mux.HandleFunc("GET /lesson23", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 23, "lesson23") }))
	//Створення обробника для шляху GET /lesson24
	mux.HandleFunc("GET /lesson24", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 24, "lesson24") }))
	//Створення обробника для шляху GET /lesson25
	mux.HandleFunc("GET /lesson25", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 25, "lesson25") }))
	//Створення обробника для шляху GET /lesson26
	mux.HandleFunc("GET /lesson26", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 26, "lesson26") }))
	//Створення обробника для шляху GET /lesson27
	mux.HandleFunc("GET /lesson27", app.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) { app.LessonHandler(w, r, 27, "lesson27") }))

	//Створення обробника для шляху GET /seal
	mux.HandleFunc("GET /seal", app.AuthMiddleware(app.SealPageHandler))

	//Сервер на https
	Logger.Info("Server launched at https://localhost:8080/")

	// Створення захисту від CSRF атаки
	CSRF := csrf.Protect(
		[]byte("32-byte-long-auth-key-goes-here!!"),
		csrf.Secure(true),
	)

	// Обгортка маршрутизатора mux у CSRF захист
	log.Fatal(http.ListenAndServeTLS(":8080", "server.crt", "server.key", loggingMiddleware(app.Logger)(CSRF(mux))))

}

// Функція автоматичного логування кожного запиту
func loggingMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("incoming request",
				"method", r.Method,
				"path", r.URL.Path,
				"remote", r.RemoteAddr,
				"agent", r.UserAgent(),
			)
			next.ServeHTTP(w, r)
		})
	}
}
