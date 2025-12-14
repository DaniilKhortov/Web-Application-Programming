package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Структура даних, що передаватиметься до клієнту (форми) у разі помилки вводу
type FormData struct {
	Username string
	Power    string
	Error    string
}

// Структура даних, що надсилатиметься до клієнту
type ConfirmationData struct {
	Username string
	Power    string
}

func main() {
	//Створення обробника для шляху /form
	http.HandleFunc("/form", formHandler)

	fmt.Println("Server launched at http://localhost:8080/form")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Функція обробки запитів від клієнта
func formHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//Обробка запиту GET для передачі даних для форми
	case http.MethodGet:
		renderForm(w, FormData{})

	//Обробка запиту POST для зміни інформації на сервері
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Form parsing error", http.StatusBadRequest)
			return
		}

		username := r.PostFormValue("username")
		powerStr := r.PostFormValue("power")

		var errorMsg string
		if username == "" {
			errorMsg = "Username field cannot be empty!"
		} else if powerStr == "" {
			errorMsg = "Power field cannot be empty!"
		} else if _, err := strconv.Atoi(powerStr); err != nil {
			errorMsg = "Power field must contain number!"
		}

		//У випадку некоректного вводу, до клієнту надсилатиметься повідомлення з помилкою
		if errorMsg != "" {

			renderForm(w, FormData{
				Username: username,
				Power:    powerStr,
				Error:    errorMsg,
			})
			return
		}
		//У ввипадку правильного вводу, перенаправлення користувача на сторінку з даними
		renderConfirmation(w, ConfirmationData{
			Username: username,
			Power:    powerStr,
		})
	//Ігнорування непередбачених запитів
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Функція надсилання помилки до клієнту
func renderForm(w http.ResponseWriter, data FormData) {
	tmpl, err := template.ParseFiles("templates/form.html")
	if err != nil {
		http.Error(w, "Template parsing error.", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

// Функція переведення на сторінку при правильних введених даних
func renderConfirmation(w http.ResponseWriter, data ConfirmationData) {
	tmpl, err := template.ParseFiles("templates/confirmation.html")
	if err != nil {
		http.Error(w, "Confirmation template parsing error!", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}
