package main

import (
	"fmt"
	"log"
	"net/http"
)

var queue []string

func main() {
	//Завантаження конфігурації
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal("Configuration error:", err)
	}

	log.Printf("Environment: %s\n", cfg.Env)
	log.Printf("Service: %s\n", cfg.ServiceName)

	//Обробка http маршруту для запису клієнта до черги
	http.HandleFunc("/enqueue", func(w http.ResponseWriter, r *http.Request) {
		if len(queue) >= cfg.MaxQueueSize {
			http.Error(w, "Queue is empty", http.StatusBadRequest)
			return
		}

		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Anonymous"
		}

		queue = append(queue, name)
		fmt.Fprintf(w, "Client %s added. Position in queue: %d\n", name, len(queue))
	})

	log.Printf("Server launched at port %s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
