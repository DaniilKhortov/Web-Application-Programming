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

	//Ініціалізація логера
	logger.Init()

	//Завантаження конфігурації черги
	cfg := config.Load()

	//Утворення черги
	var q interfaces.QueueService = queue.New(cfg.MaxSize)

	//Додавання елементів у чергу
	q.Enqueue("Task A")
	q.Enqueue("Task B")

	//Вивід інформації черги
	logger.Log.Infof("Queue size: %d", q.Size())
	color.Green("Finished!")
}
