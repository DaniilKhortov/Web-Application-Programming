package main

import (
	"fmt"
	"sync"
	"time"
)

func serveClient(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Goroutine %d starting (client in queue)\n", id)

	// Імітація обробки запиту клієнта
	time.Sleep(100 * time.Millisecond)

	fmt.Printf("Goroutine %d finished (client served)\n", id)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	for i := 1; i <= 3; i++ {
		go serveClient(i, &wg)
	}

	wg.Wait()

	fmt.Println("Everyone is served!")
}
