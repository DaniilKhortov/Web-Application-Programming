package main

import (
	"html/template"
	"log"
	"net/http"
)

// Структура для зберігання даних про сенсор
type SensorData struct {
	ID    string
	Value float64
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Парсимо HTML-шаблон із файлу
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Template parse error", http.StatusInternalServerError)
			log.Println("Template parse error:", err)
			return
		}

		// Створюємо об’єкт SensorData з тестовими даними
		data := SensorData{
			ID:    "Sensor-001",
			Value: 123.45,
		}

		// Виконуємо шаблон із передачею структури
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execute error", http.StatusInternalServerError)
			log.Println("Template execute error:", err)
		}
	})

	log.Println("Server launched at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server launch error", err)
	}
}
