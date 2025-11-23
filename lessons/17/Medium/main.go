package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Структура для JSON-відповіді
type PowerData struct {
	PowerGenerated int    `json:"power_generated"`
	Units          string `json:"units"`
}

func main() {
	// Реєстрація маршруту для кореневого шляху "/"
	http.HandleFunc("/", rootHandler)

	// Новий маршрут /status
	http.HandleFunc("/status", statusHandler)

	// Новий маршрут /data
	http.HandleFunc("/data", dataHandler)

	// Повідомлення у консолі
	fmt.Println("Server deployed at 8081...")

	// Запуск HTTP-сервера з обробкою помилок
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Error at server deployment:", err)
	}
}

// Обробник для маршруту "/"
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is running")
}

// Обробник для маршруту "/status"
func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Status: OK")
}

// Обробник для маршруту "/data"
func dataHandler(w http.ResponseWriter, r *http.Request) {
	// Створюємо об’єкт з даними
	data := PowerData{
		PowerGenerated: 1500,
		Units:          "watts",
	}

	// Встановлюємо правильний заголовок для JSON
	w.Header().Set("Content-Type", "application/json")

	// Кодуємо дані у JSON та надсилаємо клієнту
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Error generating JSON", http.StatusInternalServerError)
	}
}
