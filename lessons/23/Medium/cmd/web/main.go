package main

import (
	"log/slog"
	"net/http"
	"os"
	"queueapp/internal/handler"
)

func main() {
	// Ініціалізація структурованого логера
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Starting server", "addr", ":8080")

	// Ініціалізація маршрутизатора
	mux := http.NewServeMux()
	queueHandler := handler.NewQueueHandler(logger)

	mux.HandleFunc("/", queueHandler.ShowQueue)
	mux.HandleFunc("/add", queueHandler.AddClient)

	// Запуск сервера
	if err := http.ListenAndServe(":8080", mux); err != nil {
		logger.Error("Server failed", "error", err)
	}
}
