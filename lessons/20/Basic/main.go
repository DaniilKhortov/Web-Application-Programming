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

func main() {
	// Статичний зріз метрик
	metrics := []Metric{
		{Name: "waiting_clients", Value: 5},
		{Name: "served_clients", Value: 12},
		{Name: "average_wait_time", Value: 3.5},
	}

	// Обробник GET /metrics
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		// Перевіряємо, що запит саме GET
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Встановлюємо заголовок Content-Type
		w.Header().Set("Content-Type", "application/json")

		// Серіалізуємо зріз metrics у JSON і записуємо у відповідь
		if err := json.NewEncoder(w).Encode(metrics); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	// Запуск HTTP-сервера на порті 8080
	log.Println("Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
