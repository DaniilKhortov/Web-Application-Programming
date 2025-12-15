package main

import (
	"fmt"
	"time"
)

// Функція обробки клієнтів
func ClientProcessor(processingDone chan<- int, shutdown <-chan struct{}) {
	clientID := 1

	for {
		select {
		case <-shutdown:
			fmt.Println("Operator: shutdown command reciewed")
			return
		default:
			time.Sleep(100 * time.Millisecond)
			processingDone <- clientID
			clientID++
		}
	}
}

// Функція моніторингу статусу черги
func StatusMonitor(processingDone <-chan int, shutdown <-chan struct{}) {
	for {
		select {
		case id := <-processingDone:
			fmt.Println("Servicing client:", id)

		case <-shutdown:
			fmt.Println("StatusMonitor: finishing task...")
			return
		}
	}
}

func main() {
	// небуферизована обробка
	processingDone := make(chan int)
	shutdown := make(chan struct{})

	go ClientProcessor(processingDone, shutdown)
	go StatusMonitor(processingDone, shutdown)

	time.Sleep(1 * time.Second)

	// Глобальне завершення
	close(shutdown)

	// Невелика пауза, щоб горутини коректно завершились
	time.Sleep(200 * time.Millisecond)
	fmt.Println("E-Queue system stopped")
}
