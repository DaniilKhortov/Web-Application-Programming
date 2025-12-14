package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Структура для клієнтів черги
type QueueItem struct {
	Number int
	Name   string
}

// Обробник кореневого маршруту
func queueHandler(w http.ResponseWriter, r *http.Request) {
	queue := []QueueItem{
		{1, "Olena"},
		{2, "Dmytro"},
		{3, "Katheryna"},
		{4, "Alexij"},
	}

	//Вивід сторінки
	tmpl, err := template.ParseFiles("templates/queue.html")
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, queue)
	if err != nil {
		http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	//Встановлення маршрутизації
	http.HandleFunc("/", queueHandler)
	port := 4430
	fmt.Printf("Server is run on: https://localhost:%d\n", port)
	fmt.Println("There may appear a warning due to the self-subscribed sertificate.")

	//Шифрування за сертифікатом
	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", port), "server.crt", "server.key", nil)
	if err != nil {
		log.Fatalf("Server launch error: %v", err)
	}
}
