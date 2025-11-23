package utils

import (
	"electronic-queue/models"
	"fmt"
)

// Queue — глобальний зріз клієнтів
var Queue []models.QueueItem

// AddClient додає нового клієнта
func AddClient(name string) {
	item := models.NewQueueItem(name)
	Queue = append(Queue, item)
	fmt.Printf("Successfully added client: %s (№%d)\n", item.Client, item.ID)
}

// ShowQueue показує поточну чергу
func ShowQueue() {
	if len(Queue) == 0 {
		fmt.Println("Queue is empty.")
		return
	}
	fmt.Println("Curent queue:")
	for _, item := range Queue {
		status := "pending"
		if item.Serviced {
			status = "served"
		}
		fmt.Printf("№%d — %s (%s)\n", item.ID, item.Client, status)
	}
}

// ServeNext обслуговує наступного клієнта
func ServeNext() {
	if len(Queue) == 0 {
		fmt.Println("No one to serve!")
		return
	}
	Queue[0].Serviced = true
	fmt.Printf("Client was served: %s (#%d)\n", Queue[0].Client, Queue[0].ID)
	Queue = Queue[1:]
}
