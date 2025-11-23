package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Структура для передачі даних у шаблон форми та помилок
type FormData struct {
	Username string
	Power    string
	Error    string
}

// Структура для підтвердження
type ConfirmationData struct {
	Username string
	Power    string
}

func main() {
	http.HandleFunc("/form", formHandler)

	fmt.Println("Server launched at http://localhost:8080/form")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		renderForm(w, FormData{})

	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Form parsing error", http.StatusBadRequest)
			return
		}

		username := r.PostFormValue("username")
		powerStr := r.PostFormValue("power")

		// === Валідація ===
		var errorMsg string
		if username == "" {
			errorMsg = "Username field cannot be empty!"
		} else if powerStr == "" {
			errorMsg = "Power field cannot be empty!"
		} else if _, err := strconv.Atoi(powerStr); err != nil {
			errorMsg = "Power field must contain number!"
		}

		if errorMsg != "" {
			// Валідація не пройдена — повторно показуємо форму з помилкою
			renderForm(w, FormData{
				Username: username,
				Power:    powerStr,
				Error:    errorMsg,
			})
			return
		}

		// Дані коректні — показуємо сторінку підтвердження
		renderConfirmation(w, ConfirmationData{
			Username: username,
			Power:    powerStr,
		})

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Функція для відображення форми
func renderForm(w http.ResponseWriter, data FormData) {
	tmpl, err := template.ParseFiles("templates/form.html")
	if err != nil {
		http.Error(w, "Template parsing error.", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

// Функція для відображення підтвердження
func renderConfirmation(w http.ResponseWriter, data ConfirmationData) {
	tmpl, err := template.ParseFiles("templates/confirmation.html")
	if err != nil {
		http.Error(w, "Confirmation template parsing error!", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}
