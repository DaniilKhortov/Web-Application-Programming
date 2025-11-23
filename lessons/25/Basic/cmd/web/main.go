package main

import (
	"log"
	"net/http"
	"queueapp/internal/handler"
)

func main() {
	// Простий лог про запуск
	log.Printf("Starting server on :8080...")

	// Реєструємо обробники
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.QueueHandler)

	// Запуск сервера
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
