package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	//Реєстрація функцію-обробника для URL-шляху
	http.HandleFunc("/", rootHandler)

	fmt.Println("Server deployed at 8081...")

	// Запуск серверу на порту 8081
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Error at server deployment:", err)
	}
}

// Функція-обробник
func rootHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Server is running")
}
