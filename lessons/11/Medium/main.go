package main

import (
	"fmt"
	"sync"
	"time"
)

func serveClient(id int, wg *sync.WaitGroup, mu *sync.Mutex, results map[int]string) {
	defer wg.Done()

	fmt.Printf("Goroutine %d starting (client in queue)\n", id)

	time.Sleep(time.Duration(100+id*10) * time.Millisecond)

	mu.Lock()
	results[id] = fmt.Sprintf("client %d served at %v", id, time.Now().Format("15:04:05.000"))
	mu.Unlock()

	fmt.Printf("Goroutine %d finished (client served)\n", id)
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	results := make(map[int]string)

	numClients := 8

	wg.Add(numClients)

	for i := 1; i <= numClients; i++ {
		go serveClient(i, &wg, &mu, results)
	}

	wg.Wait()

	fmt.Println("\nEveryone is served! Results:")
	mu.Lock()
	for _, result := range results {
		fmt.Println(result)
	}
	mu.Unlock()

	fmt.Println("\nWork finished!")
}
