package services

import (
	"Medium/models"
	"Medium/utils"
	"fmt"

	"github.com/fatih/color"
)

var queue []models.QueueItem

var counter int

func init() {
	color.Blue("Initializing gueue")
	queue = []models.QueueItem{}
	counter = 0
}

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
