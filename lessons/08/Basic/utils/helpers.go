package utils

import (
	"electronic-queue/models"
	"fmt"
)

var Queue []models.QueueItem

// Функція AddClient додає клієнта до черги
func AddClient(name string) {
	//Реєстрація елементу черги
	item := models.NewQueueItem(name)
	//Додавання елементу черги
	Queue = append(Queue, item)
	fmt.Printf("Successfully added client: %s (№%d)\n", item.Client, item.ID)
}

// Функція ShowQueue ітеративно показує значення кожного елементу черги
// Або вказує, що черга порожня
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

// Функція ServeNext обслуговує першого клієнта в черзі, після чого скорочує чергу
// Або вказує, що черга порожня
func ServeNext() {
	if len(Queue) == 0 {
		fmt.Println("No one to serve!")
		return
	}
	Queue[0].Serviced = true
	fmt.Printf("Client was served: %s (#%d)\n", Queue[0].Client, Queue[0].ID)
	Queue = Queue[1:]
}
