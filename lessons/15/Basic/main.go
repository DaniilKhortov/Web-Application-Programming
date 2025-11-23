package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Реєстрація базового обробника для кореневого шляху "/"
	http.HandleFunc("/", rootHandler)

	// Інформаційне повідомлення в консолі
	fmt.Println("Server deployed at 8081...")

	// Запуск HTTP-сервера на порту 8081 з обробкою можливих помилок
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Error at server deployment:", err)
	}
}

// Обробник для кореневого маршруту "/"
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Відправляємо просте текстове повідомлення у відповідь клієнту
	fmt.Fprintln(w, "Server is running")
}
