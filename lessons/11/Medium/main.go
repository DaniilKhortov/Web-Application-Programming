package main

import (
	"fmt"
	"sync"
	"time"
)

// sensorWorker імітує датчик, який надсилає кількість клієнтів у черзі
func sensorWorker(id int, dataCh chan int, wg *sync.WaitGroup) {
	defer wg.Done() // сигналізуємо про завершення goroutine

	for i := 1; i <= 4; i++ {
		value := id*10 + i // наприклад, унікальне значення від кожного датчика
		dataCh <- value
		fmt.Printf("Counter #%d: sent value %d\n", id, value)
		time.Sleep(300 * time.Millisecond) // імітація затримки між відправленнями
	}
	fmt.Printf("Counter #%d: finished work\n", id)
}

func main() {
	// Створення буферизованого каналу розміром 5 для передачі цілих чисел
	dataCh := make(chan int, 5)

	// Використання WaitGroup для синхронізації трьох goroutines
	var wg sync.WaitGroup
	wg.Add(3)

	// Запуск трьох датчиків (goroutines)
	for i := 1; i <= 3; i++ {
		go sensorWorker(i, dataCh, &wg)
	}

	// Гороутина-агрегатор (головна) читає всі значення з каналу
	go func() {
		wg.Wait()     // чекаємо завершення всіх датчиків
		close(dataCh) // після завершення — закриваємо канал
	}()

	fmt.Println("[Agregator] Waiting for data ...")

	// Читаємо значення з каналу за допомогою for range
	for value := range dataCh {
		fmt.Printf("[Agregator] reciewed data %d\n", value)
	}

	fmt.Println("Work finished! Channel closed!")
}
