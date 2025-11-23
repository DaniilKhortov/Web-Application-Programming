package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Обробник головної сторінки
	http.Handle("/", http.FileServer(http.Dir("static")))

	// Обробник для форми /submit
	http.HandleFunc("/submit", submitHandler)

	fmt.Println("Server launched at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// submitHandler приймає POST-запит і виводить дані у консоль
func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Парсинг даних форми
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Form process error", http.StatusBadRequest)
		return
	}

	username := r.PostFormValue("username")

	// Вивід у консоль (імітація запису користувача в електронну чергу)
	fmt.Println("Recieved new user:", username)

	// Відповідь користувачу
	fmt.Fprintf(w, "Thanks, %s! You were registered into queue successfully.", username)
}
