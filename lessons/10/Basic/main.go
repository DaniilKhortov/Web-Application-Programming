package main

import (
	"fmt"
	"sync"
	"time"
)

// Функція, що імітує обслуговування клієнта
func serveClient(id int, wg *sync.WaitGroup) {
	defer wg.Done() // повідомляємо, що горутина завершилась

	fmt.Printf("Goroutine %d starting (client in queue)\n", id)

	// Імітація обробки запиту клієнта
	time.Sleep(100 * time.Millisecond)

	fmt.Printf("Goroutine %d finished (client served)\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Додаємо 3 завдання для очікування
	wg.Add(3)

	// Запускаємо 3 горутини (3 клієнти в електронній черзі)
	for i := 1; i <= 3; i++ {
		go serveClient(i, &wg)
	}

	// Очікуємо завершення всіх горутин
	wg.Wait()

	fmt.Println("Everyone is served!")
}
