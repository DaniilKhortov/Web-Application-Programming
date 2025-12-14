package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// Структура для обміну даними між клієнтом та сервером
type Metric struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

// Глобальна змінна для записів
var (
	metrics []Metric
	nextID  = 1
	mu      sync.Mutex
)

func main() {

	//Ініціалізація Mutex
	mux := http.NewServeMux()

	//Створення обробника для шляху /metrics
	mux.HandleFunc("/metrics", metricsHandler)

	//Створення обробника для шляху /metrics/
	mux.HandleFunc("/metrics/", metricHandler)

	//Додавання параметрів черги
	metrics = append(metrics, Metric{ID: nextID, Name: "waiting_clients", Value: 5})
	nextID++
	metrics = append(metrics, Metric{ID: nextID, Name: "served_clients", Value: 12})
	nextID++
	metrics = append(metrics, Metric{ID: nextID, Name: "average_wait_time", Value: 3.5})
	nextID++

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// Обробник шляху /metrics
func metricsHandler(w http.ResponseWriter, r *http.Request) {
	//Блокування даних
	mu.Lock()

	//Розблокування даних після завершення операції
	defer mu.Unlock()

	switch r.Method {

	//Обробка запиту GET
	case http.MethodGet:
		//Встановлення заголовку для надісланих повідомлень
		w.Header().Set("Content-Type", "application/json")

		//Кодування повідомлення до JSON
		json.NewEncoder(w).Encode(metrics)

	//Обробка запиту POST
	case http.MethodPost:
		var newMetric Metric
		if err := json.NewDecoder(r.Body).Decode(&newMetric); err != nil {
			http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
			return
		}
		//Додавання нових даних до черги
		newMetric.ID = nextID
		nextID++
		metrics = append(metrics, newMetric)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newMetric)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// Обробник шляху /metrics/
func metricHandler(w http.ResponseWriter, r *http.Request) {
	//Блокування даних
	mu.Lock()
	//Розблокування даних після завершення операції
	defer mu.Unlock()

	//Перевірка на правильність шляху до метрики
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.NotFound(w, r)
		return
	}

	//Конвертація ідентифікатора метрики
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	//Знаходження шляху до метрики
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
	//Обробка запиту GET для надсилання даних до клієнту
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(metrics[index])
	//Обробка запиту PUT для модифікації запису даних на сервері
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

	//Обробка запиту DELETE для видалиння запису з поля
	case http.MethodDelete:
		metrics = append(metrics[:index], metrics[index+1:]...)
		w.WriteHeader(http.StatusNoContent)
	//Ігнорування інших запитів
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
