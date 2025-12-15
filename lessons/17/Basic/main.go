package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Реєстрація обробника для кореневого маршруту "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server is running")
	})

	// Запуск HTTP-сервера на порту 8081
	fmt.Println("Starting server on http://localhost:8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
