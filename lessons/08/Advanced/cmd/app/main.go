package main

import (
	"e-queue/internal/config"
	"e-queue/internal/logger"
	"e-queue/pkg/interfaces"
	"e-queue/pkg/queue"

	"github.com/fatih/color"
)

func main() {
	color.Cyan("Running queue...")

	// Ініціалізація логера та конфігурації
	logger.Init()
	cfg := config.Load()

	// Використання інтерфейсу
	var q interfaces.QueueService = queue.New(cfg.MaxSize)
	q.Enqueue("Task A")
	q.Enqueue("Task B")

	logger.Log.Infof("Queue size: %d", q.Size())
	color.Green("Finished!")
}
