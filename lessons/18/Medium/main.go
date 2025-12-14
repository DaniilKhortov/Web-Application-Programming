package main

import (
	"fmt"
	"log"
	"net/http"
)

var queue []string

//Запуск різних середовищ:
//export GO_ENV=development
//go run .

//export GO_ENV=production
//go run .

func main() {
	//Завантаження конфігурації
	cfg := LoadConfig()

	log.Printf("Lounchung environment: %s\n", cfg.Env)
	log.Printf("Service: %s\n", cfg.ServiceName)

	//Обробка http маршруту для запису клієнта до черги
	http.HandleFunc("/enqueue", func(w http.ResponseWriter, r *http.Request) {
		if len(queue) >= cfg.MaxQueueSize {
			http.Error(w, "Queue is full", http.StatusBadRequest)
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
