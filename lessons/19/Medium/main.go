package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Структура для відповіді
type ThankYouData struct {
	Username string
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
	//Обробка запиту GET для виведення форми
	case http.MethodGet:

		//form.html є основною сторінкою
		tmpl, err := template.ParseFiles("templates/form.html")
		if err != nil {
			http.Error(w, "Form parsing error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)

	//Обробка запиту POST для зміни інформації на сервері
	case http.MethodPost:

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Form process error", http.StatusBadRequest)
			return
		}

		username := r.PostFormValue("username")
		fmt.Println("Recieved new user:", username)

		//Перенаправлення на сторінку thanks.html
		tmpl, err := template.ParseFiles("templates/thanks.html")
		if err != nil {
			http.Error(w, "Response error", http.StatusInternalServerError)
			return
		}

		//Надсилання відповіді
		data := ThankYouData{Username: username}
		tmpl.Execute(w, data)

	//Ігнорування непередбачених запитів
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
