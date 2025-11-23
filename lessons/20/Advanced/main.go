package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// Metric описує структуру метрики
type Metric struct {
	ID    int     `json:"id"`    // Унікальний ідентифікатор
	Name  string  `json:"name"`  // Назва метрики
	Value float64 `json:"value"` // Значення метрики
}

// Глобальний зріз для імітації сховища даних та м'ютекс для безпечного доступу
var (
	metrics []Metric
	nextID  = 1
	mu      sync.Mutex
)

func main() {
	mux := http.NewServeMux()

	// Маршрут для роботи з усіма метриками
	mux.HandleFunc("/metrics", metricsHandler)

	// Маршрут для роботи з конкретною метрикою по ID
	mux.HandleFunc("/metrics/", metricHandler) // обов'язково закінчується слешем

	// Початкові дані
	metrics = append(metrics, Metric{ID: nextID, Name: "waiting_clients", Value: 5})
	nextID++
	metrics = append(metrics, Metric{ID: nextID, Name: "served_clients", Value: 12})
	nextID++
	metrics = append(metrics, Metric{ID: nextID, Name: "average_wait_time", Value: 3.5})
	nextID++

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// Обробник для GET (усіх) та POST (створення) метрик
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
		w.WriteHeader(http.StatusCreated) // 201 Created
		json.NewEncoder(w).Encode(newMetric)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// Обробник для GET, PUT, DELETE конкретної метрики за ID
func metricHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Витягуємо ID з URL: /metrics/{id}
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

	// Шукаємо метрику
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
		w.WriteHeader(http.StatusNoContent) // 204 No Content
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
