package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("static")))

	http.HandleFunc("/submit", submitHandler)

	fmt.Println("Server launched at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

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
