package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Створення кореневого обробника
	http.Handle("/", http.FileServer(http.Dir("static")))

	//Створення обробника для шляху /submit
	http.HandleFunc("/submit", submitHandler)

	fmt.Println("Server launched at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Функція обробки запитів від клієнта
func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Form process error", http.StatusBadRequest)
		return
	}

	username := r.PostFormValue("username")

	fmt.Println("Recieved new user:", username)

	fmt.Fprintf(w, "Thanks, %s! You were registered into queue successfully.", username)
}
