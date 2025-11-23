package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Metric описує структуру метрики для веб-додатку
type Metric struct {
	Name  string  `json:"name"`  // Назва метрики
	Value float64 `json:"value"` // Значення метрики
}

// Глобальний зріз для імітації сховища даних
var metrics []Metric

func main() {
	// Початкові статичні метрики
	metrics = []Metric{
		{Name: "waiting_clients", Value: 5},
		{Name: "served_clients", Value: 12},
		{Name: "average_wait_time", Value: 3.5},
	}

	// Обробник для GET та POST /metrics
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// GET: повертаємо весь список метрик
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(metrics); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		case http.MethodPost:
			// POST: читаємо JSON з тіла запиту та додаємо нову метрику
			var newMetric Metric
			if err := json.NewDecoder(r.Body).Decode(&newMetric); err != nil {
				http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
				return
			}

			metrics = append(metrics, newMetric) // додаємо до глобального зрізу
			w.WriteHeader(http.StatusCreated)    // 201 Created
		default:
			// Якщо інший метод
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Запуск сервера
	log.Println("Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
