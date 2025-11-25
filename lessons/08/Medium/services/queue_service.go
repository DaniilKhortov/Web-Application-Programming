package services

import (
	"Medium/models"
	"Medium/utils"
	"fmt"

	"github.com/fatih/color"
)

// Утворення черги
var queue []models.QueueItem

// Лічильник клієнтів
var counter int

// Фуекція init ініціалізовує параметри модулю
func init() {
	color.Blue("Initializing gueue")
	queue = []models.QueueItem{}
	counter = 0
}

// Функція AddClient додає клієнта до черги
func AddClient(name string) {
	counter++
	item := models.QueueItem{
		ID:       counter,
		Client:   name,
		Serviced: false,
	}
	queue = append(queue, item)

	color.Green("Successfully added client: %s (№%d)\n", item.Client, item.ID)
	utils.PrintDivider()
}

// Функція ShowQueue ітеративно показує значення кожного елементу черги
// Або вказує, що черга порожня
func ShowQueue() {
	if len(queue) == 0 {
		color.Yellow("Queue is empty.")
		return
	}
	fmt.Println("Current queue:")
	for _, item := range queue {
		status := color.New(color.FgCyan).SprintFunc()("pending")
		if item.Serviced {
			status = color.New(color.FgMagenta).SprintFunc()("serving...")
		}
		fmt.Printf("№%d — %s (%s)\n", item.ID, item.Client, status)
	}
	utils.PrintDivider()
}

// Функція ServeNext обслуговує першого клієнта в черзі, після чого скорочує чергу
// Або вказує, що черга порожня
func ServeNext() {
	if len(queue) == 0 {
		color.Red("No one to serve")
		return
	}
	first := queue[0]
	color.Yellow("Served client: %s (№%d)", first.Client, first.ID)
	queue = queue[1:]
	utils.PrintDivider()
}
