package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Metric struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func main() {

	metrics := []Metric{
		{Name: "waiting_clients", Value: 5},
		{Name: "served_clients", Value: 12},
		{Name: "average_wait_time", Value: 3.5},
	}

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(metrics); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	log.Println("Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
