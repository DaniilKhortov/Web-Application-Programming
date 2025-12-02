package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Створення структури даних, що передаватимуться між клієнтом та сервером
type PowerData struct {
	PowerGenerated int    `json:"power_generated"`
	Units          string `json:"units"`
}

func main() {
	//Реєстрація функцій-обробників для URL-шляхів
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/data", dataHandler)

	fmt.Println("Server deployed at 8081...")

	// Запуск серверу на порту 8081
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Error at server deployment:", err)
	}
}

// Обробник кореневого маршруту /
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is running")
}

// Функція обробник для маршруту /status
func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Status: OK")
}

// Функція обробник для маршруту /data
func dataHandler(w http.ResponseWriter, r *http.Request) {

	//Створення структурованого повідомлення
	data := PowerData{
		PowerGenerated: 1500,
		Units:          "watts",
	}

	//Додавання заголовку до повідомлення
	w.Header().Set("Content-Type", "application/json")

	//Кодування повідомлення
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Error generating JSON", http.StatusInternalServerError)
	}
}
