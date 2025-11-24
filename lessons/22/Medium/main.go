package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type QueueItem struct {
	Number int
	Name   string
}

func queueHandler(w http.ResponseWriter, r *http.Request) {
	queue := []QueueItem{
		{1, "Olena"},
		{2, "Dmytro"},
		{3, "Katheryna"},
		{4, "Alexij"},
	}

	tmpl, err := template.ParseFiles("templates/queue.html")
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, queue)
	if err != nil {
		http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		tmpl, err := template.ParseFiles("templates/register.html")
		if err != nil {
			http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)

	case http.MethodPost:

		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" || password == "" {
			http.Error(w, "Fill all the fields!", http.StatusBadRequest)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Помилка хешування пароля", http.StatusInternalServerError)
			return
		}

		fmt.Println("=== Registering user ===")
		fmt.Println("Name:", username)
		fmt.Println("Hashed password:", string(hash))
		fmt.Println("=====================================")

		fmt.Fprintf(w, "Registration is successful!")
	default:
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", queueHandler)
	http.HandleFunc("/register", registerHandler)

	port := 4430
	fmt.Printf("Server is run on: https://localhost:%d\n", port)
	fmt.Println("There may appear a warning due to the self-subscribed sertificate.")

	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", port), "server.crt", "server.key", nil)
	if err != nil {
		log.Fatalf("Server launch error: %v", err)
	}
}
