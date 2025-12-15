package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// поточний номер у черзі
var currentNumber int = 0

// кількість обслугованих
var servedClients int = 0

// м'ютекс для безпечного доступу
var mu sync.Mutex

// Структура для JSON-відповіді
type MetricsResponse struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

func main() {
	mux := http.NewServeMux()

	// Стандартні маршрути
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server is running")
	})

	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Status: OK")
	})

	// Інкремент наступного номера
	mux.HandleFunc("/next", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		currentNumber++
		num := currentNumber
		mu.Unlock()

		response := map[string]int{"next_number": num}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Просунутий маршрут /api/metrics з параметром type
	mux.HandleFunc("/api/metrics", func(w http.ResponseWriter, r *http.Request) {
		metricType := r.URL.Query().Get("type")
		var value interface{}

		mu.Lock()
		switch metricType {
		case "queue":
			value = currentNumber
		case "served":
			value = servedClients
		case "waiting":
			value = currentNumber - servedClients
		default:
			value = "unknown type"
		}
		mu.Unlock()

		response := MetricsResponse{
			Type:  metricType,
			Value: value,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Запуск HTTP-сервера на порту 8081
	fmt.Println("Starting server on http://localhost:8081")
	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
