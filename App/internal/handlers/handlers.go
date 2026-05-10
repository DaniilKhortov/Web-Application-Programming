package handler

import (
	"QueueApp/internal/models"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"text/template"

	"time"

	"github.com/gorilla/csrf"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Структура App для зберігання конектора БД, логера і вмісту сторінок
type App struct {
	Logger *slog.Logger
	Data   map[int]models.Lesson
	DB     *gorm.DB
}

// Данні для роботи з gitLab (API)
const (
	gitlabToken = "glpat-ih0DDOm0toKeXi-PfNLsJGM6MQpvOjEKdTptbDQwbA8.01.171s2sr77"
	projectID   = "81919874"
	gitlabURL   = "https://gitlab.com/api/v4/projects"
)

// Функція логування помилки і надсилання 500 Internal Server Error
func (a *App) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)
	a.Logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Функція відправки стану клієнту
func (a *App) clientError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

// Обробник шляху /index.html
func (a *App) HomeHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		a.serverError(w, r, err)
		return
	}
	tmpl.Execute(w, nil)
}

// Обробник шляху /landing.html
func (a *App) LandingHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		a.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("web/templates/landing.html")
	if err != nil {
		a.serverError(w, r, err)
		return
	}
	tmpl.Execute(w, nil)
}

// Обробник шляху /lesson1.html
func (a *App) LessonHandler(w http.ResponseWriter, r *http.Request, number int, template string) {
	lesson, ok := a.Data[number]
	if !ok {
		a.Logger.Error("Lesson not found", "id", 1)
		http.NotFound(w, r)
		return
	}
	lesson.CSRFToken = csrf.Token(r)
	a.render(w, r, "lesson1", lesson)
}

// Обробник шляху /admin.html
func (a *App) AdminPageHandler(w http.ResponseWriter, r *http.Request) {
	var files []models.FileData
	var users []models.UserData
	result := a.DB.Find(&files)
	resultU := a.DB.Find(&users)

	if result.Error != nil && resultU != nil {
		a.Logger.Error("Помилка БД: " + result.Error.Error())
		http.Error(w, "Внутрішня помилка сервера", http.StatusInternalServerError)
		return
	}
	data := map[string]interface{}{
		"FileRecords": files,
		"UserRecords": users,
		"CSRFToken":   csrf.Token(r),
	}

	a.render(w, r, "admin", data)
}

// Обробник для оновлення оцінки
func (a *App) UpdateMarkHandler(w http.ResponseWriter, r *http.Request) {

	//Імпорт id з запиту
	id := r.PathValue("id")

	var req models.FileData

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	//Пошук і оновлення запису
	result := a.DB.Model(&models.FileData{}).Where("id = ?", id).Update("mark", req.Mark)

	if result.Error != nil {
		a.Logger.Error("Помилка оновлення оцінки: " + result.Error.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Обробник для оновлення ролі користувача
func (a *App) UpdateRoleHandler(w http.ResponseWriter, r *http.Request) {
	//Імпорт id з запиту
	id := r.PathValue("id")

	var req models.UserData

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Пошук та оновлення ролі
	result := a.DB.Model(&models.UserData{}).Where("id = ?", id).Update("role", req.Role)
	if result.Error != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Оброблює шлях /register.html
func (a *App) RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("web/templates/register.html")

	// Сертифікат csrf
	data := map[string]interface{}{
		"csrfField": csrf.TemplateField(r),
	}
	tmpl.Execute(w, data)
}

// Оброблює шлях /seal.html
func (a *App) SealPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("web/templates/seal.html")

	// Сертифікат csrf
	data := map[string]interface{}{
		"csrfField": csrf.TemplateField(r),
	}
	tmpl.Execute(w, data)
}

// Оброблює шлях /login.html для перенаправлення на сторінку
func (a *App) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("web/templates/login.html")

	// Сертифікат csrf
	data := map[string]interface{}{
		"csrfField": csrf.TemplateField(r),
	}

	tmpl.Execute(w, data)
}

// Оброблює шлях /login.html для входу
func (a *App) LoginSubmitHandler(w http.ResponseWriter, r *http.Request) {
	// Перевірка методу
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Отримання даних зі сторінки
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	var u models.UserData

	// Пошук користувача через GORM у БД
	result := a.DB.Where("email = ?", email).First(&u)

	if result.Error != nil {
		a.Logger.Warn("Спроба входу: користувача не знайдено", "email", email)
		a.renderLoginError(w, r, "Невірний email або пароль!")
		return
	}

	// Перевірка хешу паролю
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		a.Logger.Warn("Спроба входу: невірний пароль", "user", u.Username)
		a.renderLoginError(w, r, "Невірний email або пароль!")
		return
	}

	// Встановлення Cookies
	a.setUserCookies(w, u.Username, u.ID)

	// Логіка перенаправлення залежно від ролі
	a.Logger.Info("Користувач увійшов", "user", u.Username, "role", u.Role)

	if u.Role == "admin" {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}
	if u.Role == "seal" {
		http.Redirect(w, r, "/seal", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/main", http.StatusSeeOther)
}

// Функція обробки реєстрації користувача
func (a *App) RegisterSubmitHandler(w http.ResponseWriter, r *http.Request) {

	//Перевірка вмісту форми
	if err := r.ParseForm(); err != nil {
		a.Logger.Error("Помилка парсингу форми", "error", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Отримання даних зі сторінки
	username := r.PostFormValue("username")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	passwordConfirm := r.PostFormValue("passwordConfirm")

	errMsg := validateUserRegister(username, email, password, passwordConfirm)
	if errMsg != "" {
		a.renderRegisterError(w, r, errMsg)
		return
	}

	//Створення хешу паролю
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		a.Logger.Error("Помилка хешування", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	//Структура користувача для зберігання в БД як запис
	newUser := models.UserData{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
		Role:     "user",
	}

	// Збереження в БД через GORM
	result := a.DB.Create(&newUser)
	if result.Error != nil {
		a.Logger.Error("Помилка створення користувача", "error", result.Error)
		a.renderRegisterError(w, r, "Користувач з таким ім'ям або email вже існує")
		return
	}

	// Встановлення Cookies
	a.setUserCookies(w, newUser.Username, newUser.ID)

	// 6. Перенаправлення на головну
	http.Redirect(w, r, "/main", http.StatusSeeOther)
}

// Оброблює шлях /login.html для входу
func (a *App) UploadHandler(w http.ResponseWriter, r *http.Request) {

	//Надання CORS-прав (мабуть буде усунуто через надлишковість)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Обробка попереднього запиту
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Error(w, "Ви не авторизовані", http.StatusUnauthorized)
		return
	}

	// Взяття даних з cookies
	userID := cookie.Value

	// Взяття даних з сторінки
	file, header, _ := r.FormFile("file")
	lessonID := r.FormValue("lessonID")

	defer file.Close()
	// Зчитування взятих даних
	content, _ := io.ReadAll(file)
	a.Logger.Info("Спроба завантажити", "Path: ", fmt.Sprintf("%s/%s/%s", lessonID, userID, header.Filename))

	// Готування шляху збереження: /lessonID/user1/назва_файлу
	filePath := fmt.Sprintf("%s/%s/%s", lessonID, userID, header.Filename)

	// Формування запиту до GitLab API
	payload := map[string]string{
		"branch":         "main",
		"author_email":   "student@example.com",
		"author_name":    "Student Bot",
		"content":        base64.StdEncoding.EncodeToString(content),
		"commit_message": "Upload work for " + userID,
		"encoding":       "base64",
	}

	jsonData, _ := json.Marshal(payload)
	encodedPath := url.PathEscape(filePath)

	// URL для створення файлу: /projects/:id/repository/files/:file_path
	apiURL := fmt.Sprintf("%s/%s/repository/files/%s", gitlabURL, projectID, encodedPath)

	req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	req.Header.Set("PRIVATE-TOKEN", gitlabToken)
	req.Header.Set("Content-Type", "application/json")

	// Відправка на GitLab
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Помилка запиту до GitLab:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("GitLab повернув помилку %d: %s\n", resp.StatusCode, string(body))
		http.Error(w, "GitLab відмовив у доступі", resp.StatusCode)
		return
	}
	now := time.Now()

	//Структура для збереження запису про завантажений ффайл у БД
	newFile := models.FileData{
		Username: userID,
		Lesson:   lessonID,
		FilePath: fmt.Sprintf("https://gitlab.com/daniilkhortov1/UserWorks/-/tree/main/%s/%s", lessonID, userID),
		Date:     now.Format("2006-01-02"),
		Mark:     "unmarked",
	}

	// Збереження запису в БД через GORM
	result := a.DB.Create(&newFile)
	if result.Error != nil {
		a.Logger.Error("Помилка створення користувача", "error", result.Error)
		// Можна додати перевірку на дублікат імені/email
		a.renderRegisterError(w, r, "Користувач з таким ім'ям або email вже існує")
		return
	}
}

// Обробник шляху /taskStatus
// Відповідає на polling клієнта про дані оцінки
func (a *App) TaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	//Взяття даних з cookie
	userID := cookie.Value
	//Взяття даних з URL
	lessonID := r.URL.Query().Get("lessonID")

	//Знаходження у БД запису роботи
	var fileData models.FileData

	result := a.DB.Where("username = ? AND lesson = ?", userID, lessonID).First(&fileData)

	w.Header().Set("Content-Type", "application/json")

	//Реакція серверу якщо роботи немає
	if result.Error != nil {
		json.NewEncoder(w).Encode(map[string]string{"status": "no_record"})
		return
	}
	//Реакція серверу якщо робота не оцінена
	if fileData.Mark == "unmarked" {
		json.NewEncoder(w).Encode(map[string]string{"status": "pending"})
	} else {
		//Реакція серверу якщо робота  оцінена
		json.NewEncoder(w).Encode(map[string]string{
			"status": "graded",
			"mark":   fileData.Mark,
		})
	}
}

// Функція валідації полів
func validateUserRegister(username, email, password, confirm string) string {
	if username == "" || email == "" || password == "" || confirm == "" {
		return "Усі поля мають бути заповнені!"
	}
	if password != confirm {
		return "Паролі не збігаються!"
	}
	if len(password) < 6 {
		return "Пароль занадто короткий (мін. 6 символів)!"
	}
	// Тут можна додати перевірку формату email через regexp
	return ""
}

// Допоміжна функція для рендеру помилок
func (a *App) renderRegisterError(w http.ResponseWriter, r *http.Request, msg string) {
	tmpl, _ := template.ParseFiles("web/templates/register.html")
	tmpl.Execute(w, map[string]interface{}{
		"Error":     msg,
		"csrfField": csrf.TemplateField(r),
	})
}

// Допоміжна функція для встановлення Cookies
func (a *App) setUserCookies(w http.ResponseWriter, username string, id uint) {
	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Value:    username,
		Path:     "/",
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "user_id",
		Value: fmt.Sprintf("%d", id),
		Path:  "/",
	})
}

// Допоміжна функція для рендеру помилок логіну
func (a *App) renderLoginError(w http.ResponseWriter, r *http.Request, msg string) {
	tmpl, err := template.ParseFiles("web/templates/login.html")
	if err != nil {
		a.serverError(w, r, err)
		return
	}
	tmpl.Execute(w, map[string]interface{}{
		"Error":     msg,
		"csrfField": csrf.TemplateField(r),
	})
}

// Middleware для унеможливлення переходу на сторінки неавторизованим
func (a *App) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Отримання Cookie
		// Якщо  немає - перенаправлення на головну
		_, err := r.Cookie("user_id")
		if err != nil {

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	}
}

// Middleware для унеможливлення переходу на сторінки персоналу
func (a *App) AdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Отримання Cookie
		// Якщо  немає - перенаправлення на головну
		cookie, err := r.Cookie("user_id")

		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// Ідентифікація користувача в БД
		var user models.UserData
		if err := a.DB.First(&user, cookie.Value).Error; err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// Перевірка ролі
		if user.Role != "admin" {
			// Якщо не адмін — відправляємо на головну
			a.Logger.Warn("Спроба несанкціонованого доступу", "user", user.Username)
			http.Redirect(w, r, "/main", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	}
}
