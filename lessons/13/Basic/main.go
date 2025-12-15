package main

import (
	"fmt"
	"sync"
)

var (
	queueCounter int
	mutex        sync.Mutex
	wg           sync.WaitGroup
)

// Функція небезпечного доступу до даних черги
func unsafeClient() {
	defer wg.Done()
	queueCounter++
}

// Функція безпечного доступу до даних черги з Mutex
func safeClient() {
	defer wg.Done()

	mutex.Lock()
	defer mutex.Unlock()

	queueCounter++
}

func main() {
	clients := 1000

	//Демонстрація гонки даних
	queueCounter = 0
	wg.Add(clients)

	for i := 0; i < clients; i++ {
		go unsafeClient()
	}

	wg.Wait()
	fmt.Println("Amount of serviced clients in queue without synchronization (Data Race):", queueCounter)
	fmt.Println("Expected: 1000")
	//Демонстрація безпечної обробки даних
	queueCounter = 0
	wg.Add(clients)

	for i := 0; i < clients; i++ {
		go safeClient()
	}

	wg.Wait()
	fmt.Println("Amount of serviced clients in queue with Mutex:", queueCounter)
	fmt.Println("Expected: 1000")
}
