// file: main.go
package main

import (
	"fmt"
	"sync"
	"time"
)

// Кількість горутин (конкурентних клієнтів)
const workers = 1000

// Імітуємо веб-додаток "електронна черга": простий загальний лічильник обслугованих клієнтів.
// Версія A: без синхронізації (демонстрація Data Race).
// Версія B: з sync.Mutex (виправлення гонки).

func simulateWithoutMutex() int {
	var counter int // shared resource — тут відбудеться гонка
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func(id int) {
			// невелика штучна затримка, щоб підсилити ймовірність гонки
			time.Sleep(time.Microsecond * time.Duration(id%5))
			// Критична секція: інкремент без захисту -> data race
			counter = counter + 1
			wg.Done()
		}(i)
	}

	wg.Wait()
	return counter
}

func simulateWithMutex() int {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func(id int) {
			// штучна затримка — як і раніше
			time.Sleep(time.Microsecond * time.Duration(id%5))
			// Захищаємо доступ до shared resource
			mu.Lock()
			defer mu.Unlock() // гарантуємо розблокування, навіть якщо буде паніка (хоча тут її немає)
			counter = counter + 1
			wg.Done()
		}(i)
	}

	wg.Wait()
	return counter
}

func main() {

	fmt.Printf("Running %d goroutines to serve clients simulteniosly.\n\n", workers)

	// 1) Демонстрація гонки даних
	fmt.Println("Data Race without Mutex:")
	noMutexCount := simulateWithoutMutex()
	fmt.Printf("Result: %d (expected %d)\n", noMutexCount, workers)
	if noMutexCount != workers {
		fmt.Println("Data race happened!")
	} else {
		fmt.Println("Somehow data race was avoided. Run again to be sure!")
	}
	fmt.Println()

	// 2) Виправлення з використанням sync.Mutex
	fmt.Println("Data Race with Mutex:")
	withMutexCount := simulateWithMutex()
	fmt.Printf("Result: %d (expected %d)\n", withMutexCount, workers)
	if withMutexCount == workers {
		fmt.Println("Data race was avoided!")
	} else {
		fmt.Println("Error! Something went wrong!")
	}
}
