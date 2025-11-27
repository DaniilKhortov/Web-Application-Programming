package main

import (
	"fmt"
	"sync"
	"time"
)

// Горутина обслуговування клієнта
func serveClient(id int, wg *sync.WaitGroup, mu *sync.Mutex, results map[int]string) {
	//Створення defer для оновлення лічильника після завершення роботи горутини
	defer wg.Done()

	fmt.Printf("Goroutine %d starting (client in queue)\n", id)

	//Імітація обробки запиту клієнта
	time.Sleep(time.Duration(100+id*10) * time.Millisecond)

	//Застосування Mutex для безпесної обробки даних
	//Mutex блокує ділянку пам'ятті від інших процесів до поки горутина не завершить модифікацію
	mu.Lock()
	results[id] = fmt.Sprintf("client %d served at %v", id, time.Now().Format("15:04:05.000"))

	//Розблокувіння ділянки пам'ятті
	mu.Unlock()

	fmt.Printf("Goroutine %d finished (client served)\n", id)
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	//Створення загальної мапи результати
	results := make(map[int]string)

	//Число клієнтів - визначає кількість горутин
	numClients := 8

	wg.Add(numClients)

	//Запуск горутини
	for i := 1; i <= numClients; i++ {
		go serveClient(i, &wg, &mu, results)
	}
	//Блокування усіх потоків до виконання усіх горутин
	wg.Wait()

	fmt.Println("\nEveryone is served! Results:")

	//Безпечний вивід результатів
	mu.Lock()
	for _, result := range results {
		fmt.Println(result)
	}
	mu.Unlock()

	fmt.Println("\nWork finished!")
}
