package main

import (
	"html/template"
	"log"
	"net/http"
)

type SensorData struct {
	ID    string
	Value float64
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Template parse error", http.StatusInternalServerError)
			log.Println("Template parse error:", err)
			return
		}

		data := SensorData{
			ID:    "Sensor-001",
			Value: 123.45,
		}

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
