package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"queueapp/internal/handler"
	"queueapp/internal/service"
	"queueapp/internal/storage"
)

func main() {
	port := flag.String("port", "8080", "HTTP server port")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	store := storage.NewMemoryStore()
	svc := service.NewQueueService(store)
	h := handler.NewQueueHandler(svc, logger)

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("/", h.Home)
	mux.HandleFunc("/create", h.Create)
	mux.HandleFunc("/edit", h.Edit)
	mux.HandleFunc("/delete", h.Delete)

	logger.Info("Starting server...", "port", *port)
	err := http.ListenAndServe(":"+*port, mux)
	if err != nil {
		logger.Error("Server failed", "error", err)
		os.Exit(1)
	}
}
