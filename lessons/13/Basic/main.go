package main

import (
	"fmt"
	"strings"
	"time"
)

// Структура для зберігання даних клієнта
type Client struct {
	Name      string
	TicketNum int
}

// --- Етап 1: DataReader ---
func DataReader(out chan<- Client) {
	// Умовно вхідні дані (можна уявити, що вони прийшли з веб-форми)
	input := []Client{
		{"Michael", 1},
		{"Maria", 2},
		{"", 3},        // некоректне ім’я
		{"Dmytro", -1}, // некоректний номер
		{"Daryna", 4},
	}

	for _, client := range input {
		fmt.Println("[Reader] Recieved record:", client)
		out <- client
		time.Sleep(300 * time.Millisecond) // імітація часу обробки
	}
	close(out)
}

// --- Етап 2: DataProcessor ---
func DataProcessor(in <-chan Client) {
	for client := range in {
		if validate(client) {
			fmt.Printf("[Processor] Client '%s' is registered as %d\n",
				client.Name, client.TicketNum)
		} else {
			fmt.Printf("[Processor] Invalid data: %+v\n", client)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

// Функція валідації вхідних даних
func validate(c Client) bool {
	if strings.TrimSpace(c.Name) == "" {
		return false
	}
	if c.TicketNum <= 0 {
		return false
	}
	return true
}

// --- Основна функція ---
func main() {
	fmt.Println("E-Queue")

	// Канал для зв’язку між Reader та Processor
	dataChan := make(chan Client)

	// Запуск етапів у окремих goroutines
	go DataReader(dataChan)
	DataProcessor(dataChan) // головна горутина чекає завершення обробки

	fmt.Println("Work is done!")
}
