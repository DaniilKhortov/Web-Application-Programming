package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Структура для JSON-відповіді (використовується в /data)
type PowerData struct {
	PowerGenerated int    `json:"power_generated"`
	Units          string `json:"units"`
}

func main() {
	// 🔹 Створюємо власний маршрутизатор (а не стандартний DefaultServeMux)
	mux := http.NewServeMux()

	// Реєструємо всі маршрути у власному mux
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/status", statusHandler)
	mux.HandleFunc("/data", dataHandler)
	mux.HandleFunc("/api/metrics", metricsHandler)

	// Повідомлення у консолі
	fmt.Println("Server deployed at 8081...")

	// Запуск сервера з власним маршрутизатором
	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal("Error at server deployment:", err)
	}
}

// ------------------- Обробники -------------------

// "/" — базовий обробник
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is running")
}

// "/status" — повертає текстовий статус
func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Status: OK")
}

// "/data" — повертає JSON-відповідь
func dataHandler(w http.ResponseWriter, r *http.Request) {
	data := PowerData{
		PowerGenerated: 1500,
		Units:          "watts",
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Error generating JSON", http.StatusInternalServerError)
	}
}

// "/api/metrics" — аналізує query-параметри, наприклад ?sensor=temp
func metricsHandler(w http.ResponseWriter, r *http.Request) {
	// Зчитуємо значення параметра sensor
	sensor := r.URL.Query().Get("sensor")

	// Якщо параметр не вказано
	if sensor == "" {
		http.Error(w, "Parameter 'sensor' is required, e.g. /api/metrics?sensor=temp", http.StatusBadRequest)
		return
	}

	// Динамічна відповідь залежно від sensor
	response := fmt.Sprintf("Metrics for sensor: %s", sensor)

	// Повертаємо динамічну відповідь
	fmt.Fprintln(w, response)
}
