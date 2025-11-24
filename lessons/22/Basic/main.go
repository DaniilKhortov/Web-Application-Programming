package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type QueueItem struct {
	Number int
	Name   string
}

func queueHandler(w http.ResponseWriter, r *http.Request) {
	queue := []QueueItem{
		{1, "Olena"},
		{2, "Dmytro"},
		{3, "Katheryna"},
		{4, "Alexij"},
	}

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
	http.HandleFunc("/", queueHandler)

	port := 4430
	fmt.Printf("Server is run on: https://localhost:%d\n", port)
	fmt.Println("There may appear a warning due to the self-subscribed sertificate.")

	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", port), "server.crt", "server.key", nil)
	if err != nil {
		log.Fatalf("Server launch error: %v", err)
	}
}
