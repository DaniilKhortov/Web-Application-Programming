package main

import (
	"fmt"
	"sync"
	"time"
)

func sender(id int, dataCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 4; i++ {
		value := id*10 + i
		fmt.Printf("Goroutine %d: adding client %d\n", id, value)
		dataCh <- value
		time.Sleep(500 * time.Millisecond) // імітація роботи
	}
}

func main() {
	// Буферизований канал розміром 5
	dataCh := make(chan int, 5)

	var wg sync.WaitGroup

	// Запуск 3 горутин для реєстрації клієнтів
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go sender(i, dataCh, &wg)
	}

	// Закриття каналу після завершення всіх відправників
	go func() {
		wg.Wait()
		close(dataCh)
		fmt.Println("All goroutines finished their task. Channel closed!")
	}()

	// Приймання значень через for range
	for value := range dataCh {
		fmt.Println("Main queue: servicing client ", value)
	}

	fmt.Println("That is all for today!")
}
