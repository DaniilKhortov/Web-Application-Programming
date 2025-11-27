package main

import (
	"fmt"
	"sync"
	"time"
)

// Горутина обслуговування клієнта
func serveClient(id int, wg *sync.WaitGroup) {
	//Створення defer для оновлення лічильника після завершення роботи горутини
	defer wg.Done()

	fmt.Printf("Goroutine %d starting (client in queue)\n", id)

	//Імітація обробки запиту клієнта
	time.Sleep(100 * time.Millisecond)

	fmt.Printf("Goroutine %d finished (client served)\n", id)
}

func main() {

	//Створення WaitGroup
	//WaitGroup очікує завершення усіх горутин
	var wg sync.WaitGroup

	//Задання лічильника горутин
	wg.Add(3)

	//Запуск горутини
	for i := 1; i <= 3; i++ {
		go serveClient(i, &wg)
	}

	//Блокування усіх потоків до виконання усіх горутин
	wg.Wait()

	fmt.Println("Everyone is served!")
}
