package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Структура для JSON-відповіді на маршруті /queue
type QueueResponse struct {
	NextNumber int `json:"next_number"`
}

// Поточний номер клієнта у черзі
var currentNumber int = 0
var mu sync.Mutex

func main() {
	// Обробник кореневого маршруту "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server is running")
	})

	// Обробник маршруту "/status"
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Status: OK")
	})

	// Обробник маршруту "/data"
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"service": "customer queue",
			"units":   "clients",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Новий маршрут "/next" для отримання наступного номера у черзі
	http.HandleFunc("/next", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		currentNumber++
		response := QueueResponse{NextNumber: currentNumber}
		mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Запуск HTTP-сервера на порту 8081
	fmt.Println("Starting server on http://localhost:8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
