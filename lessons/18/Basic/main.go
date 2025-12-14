package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

var queue []string

// Для введення клієнта в чергу, використовується адреса:
// http://localhost:8080/enqueue?name=Daniel
func main() {
	// Завантаження .env
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found! Using default parameters")
	}

	// Завантаження конфігурації
	cfg := LoadConfig()

	fmt.Println("Service:", cfg.ServiceName)
	fmt.Println("Max queue length:", cfg.MaxQueueSize)

	// HTTP обробник для додавання клієнта в чергу
	http.HandleFunc("/enqueue", func(w http.ResponseWriter, r *http.Request) {
		if len(queue) >= cfg.MaxQueueSize {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Queue is full")
			return
		}

		client := r.URL.Query().Get("name")
		if client == "" {
			client = "Anonymous"
		}

		queue = append(queue, client)
		fmt.Fprintf(w, "Client %s added. Position in queue: %d\n", client, len(queue))
	})

	log.Printf("Server deployed at: %s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
