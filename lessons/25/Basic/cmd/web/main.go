package main

import (
	"log"
	"net/http"
	"queueapp/internal/handler"
)

func main() {

	log.Printf("Starting server on :8080...")

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.QueueHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
