package main

import (
	"html/template"
	"log"
	"net/http"
)

// Структура для даних одного сенсора
type SensorData struct {
	ID    string
	Value float64
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Парсимо шаблон
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Template parse error", http.StatusInternalServerError)
			log.Println("Template parse error:", err)
			return
		}

		// Створюємо зріз із кількох датчиків
		sensors := []SensorData{
			{ID: "Sensor-001", Value: 85.2},
			{ID: "Sensor-002", Value: 102.7},
			{ID: "Sensor-003", Value: 76.4},
			{ID: "Sensor-004", Value: 145.9},
		}

		// Передаємо зріз у шаблон
		err = tmpl.Execute(w, sensors)
		if err != nil {
			http.Error(w, "Template execute error", http.StatusInternalServerError)
			log.Println("Template execute error:", err)
		}
	})

	log.Println("Server launched at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server launch error:", err)
	}
}
