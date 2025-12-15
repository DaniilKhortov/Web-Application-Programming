package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"queue-app/internal/config"
)

// Інформація про збірку
var (
	Version   = "dev"
	BuildTime = "unknown"
	GitCommit = "unknown"
)

var queue = 0

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	//Вивід збірки
	log.Printf(
		"%s started | env=%s | version=%s | build=%s | commit=%s",
		cfg.AppName, cfg.Env, Version, BuildTime, GitCommit,
	)

	//Обробник кореневого марщруту
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Електронна черга\nПоточний номер: %d\n", queue)
	})

	//Обробник шляху маршрутизації /next для додавання клієнтів
	mux.HandleFunc("/next", func(w http.ResponseWriter, _ *http.Request) {
		queue++
		fmt.Fprintf(w, "Новий номер у черзі: %d\n", queue)
	})

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: mux,
	}

	// Реалізація Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill) // SIGINT, SIGTERM
	go func() {
		<-stop
		log.Println("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Server Shutdown Failed:%+v", err)
		}
		log.Println("Server exited gracefully")
	}()

	log.Printf("Listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe(): %v", err)
	}
}
