package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Metric struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

var (
	metrics []Metric
	nextID  = 1
	mu      sync.Mutex
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/metrics", metricsHandler)

	mux.HandleFunc("/metrics/", metricHandler)

	metrics = append(metrics, Metric{ID: nextID, Name: "waiting_clients", Value: 5})
	nextID++
	metrics = append(metrics, Metric{ID: nextID, Name: "served_clients", Value: 12})
	nextID++
	metrics = append(metrics, Metric{ID: nextID, Name: "average_wait_time", Value: 3.5})
	nextID++

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(metrics)
	case http.MethodPost:
		var newMetric Metric
		if err := json.NewDecoder(r.Body).Decode(&newMetric); err != nil {
			http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
			return
		}
		newMetric.ID = nextID
		nextID++
		metrics = append(metrics, newMetric)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newMetric)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func metricHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.NotFound(w, r)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	index := -1
	for i, m := range metrics {
		if m.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		http.Error(w, "Metric not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(metrics[index])
	case http.MethodPut:
		var updated Metric
		if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
			http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
			return
		}
		updated.ID = id
		metrics[index] = updated
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updated)
	case http.MethodDelete:
		metrics = append(metrics[:index], metrics[index+1:]...)
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
