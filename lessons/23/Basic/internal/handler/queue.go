package handler

import (
	"html/template"
	"log"
	"net/http"
)

// Дані для черги (імітація бази даних)
var clients = []string{
	"Ivan Petrenko",
	"Maria Kondratenko",
	"Oleh Stupa",
	"Volodymyr Kononenko",
}

// QueueHandler — обробник, який показує HTML-сторінку зі списком клієнтів
func QueueHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request for %s", r.Method, r.URL.Path)

	tmpl, err := template.ParseFiles("web/templates/queue.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		return
	}

	data := struct {
		Title   string
		Clients []string
	}{
		Title:   "E-Queue",
		Clients: clients,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
