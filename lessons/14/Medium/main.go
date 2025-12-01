package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Структура клієнта
type Client struct {
	ID        int
	Name      string
	TicketNum int
}

var (
	ProcessedIDs []int
	mu           sync.Mutex
)

// DataGenerator - записує клієнтів до каналу
func DataGenerator(out chan<- Client) {
	input := []Client{
		{1, "Michael", 1},
		{2, "Maria", 2},
		{3, "", 3},
		{4, "Dmytro", -1},
		{5, "Daryna", 4},
	}

	for _, client := range input {
		fmt.Printf("[Generator] Client sent: %+v\n", client)
		out <- client
		time.Sleep(200 * time.Millisecond)
	}
	close(out)
}

// ParallelFilter - зчитує клієнтів з каналу та виводить його дані
// Використовує безпечну обробку данних
func ParallelFilter(id int, in <-chan Client, wg *sync.WaitGroup) {
	defer wg.Done()

	for client := range in {
		if validate(client) {
			mu.Lock()
			ProcessedIDs = append(ProcessedIDs, client.ID)
			mu.Unlock()
			fmt.Printf("[Worker %d] Client '%s' is registered as %d\n",
				id, client.Name, client.TicketNum)
		} else {
			fmt.Printf("[Worker %d] Invalid data: %+v\n", id, client)
		}
		time.Sleep(300 * time.Millisecond)
	}
}

// validate - перевіряє коректність даних
func validate(c Client) bool {
	if strings.TrimSpace(c.Name) == "" {
		return false
	}
	if c.TicketNum <= 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println("E-Queue")

	dataChan := make(chan Client)
	var wg sync.WaitGroup

	//Запуск горутини
	go DataGenerator(dataChan)

	//Обробка даних з каналу в кілька потоків
	numWorkers := 3
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go ParallelFilter(i, dataChan, &wg)
	}

	wg.Wait()

	//Вивід результатів
	fmt.Println("\nAll client served!")
	mu.Lock()
	fmt.Println("ID of clients after check:", ProcessedIDs)
	mu.Unlock()
}
