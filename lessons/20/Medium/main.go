package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Структура для обміну даними між клієнтом та сервером
type Metric struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

var metrics []Metric

func main() {

	//Створення  черги
	metrics = []Metric{
		{Name: "waiting_clients", Value: 5},
		{Name: "served_clients", Value: 12},
		{Name: "average_wait_time", Value: 3.5},
	}

	//Створення обробника для шляху /metrics
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		//Обробка запиту GET
		case http.MethodGet:
			//Встановлення заголовку для надісланих повідомлень
			w.Header().Set("Content-Type", "application/json")

			//Кодування повідомлення до JSON
			if err := json.NewEncoder(w).Encode(metrics); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

		//Обробка запиту POST
		case http.MethodPost:

			var newMetric Metric
			if err := json.NewDecoder(r.Body).Decode(&newMetric); err != nil {
				http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
				return
			}
			//Додавання нових даних до черги
			metrics = append(metrics, newMetric)
			w.WriteHeader(http.StatusCreated)
		default:

			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	//Запуск серверу на порті 8080
	log.Println("Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
