package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// структура для передачі даних у шаблон
type ThankYouData struct {
	Username string
}

func main() {
	// Обробник GET — відображення форми
	http.HandleFunc("/form", formHandler)

	fmt.Println("Server launched at http://localhost:8080/form")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// formHandler обробляє і GET, і POST на /form
func formHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Відображаємо HTML-форму
		tmpl, err := template.ParseFiles("templates/form.html")
		if err != nil {
			http.Error(w, "Form parsing error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)

	case http.MethodPost:
		// Обробляємо надіслані дані
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Form process error", http.StatusBadRequest)
			return
		}

		username := r.PostFormValue("username")
		fmt.Println("Recieved new user:", username)

		// Формуємо персоналізовану відповідь через шаблон
		tmpl, err := template.ParseFiles("templates/thanks.html")
		if err != nil {
			http.Error(w, "Response error", http.StatusInternalServerError)
			return
		}

		data := ThankYouData{Username: username}
		tmpl.Execute(w, data)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
