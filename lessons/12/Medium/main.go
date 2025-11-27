package main

import (
	"fmt"
	"sync"
	"time"
)

// Горутина, що надсилає дані у канал
func sensorWorker(id int, dataCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	//Ітеративно заповнюємо буферизований канал
	for i := 1; i <= 4; i++ {
		value := id*10 + i
		dataCh <- value
		fmt.Printf("Counter #%d: sent value %d\n", id, value)
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Printf("Counter #%d: finished work\n", id)
}

func main() {

	//Ініціалізація буфериованого каналу
	dataCh := make(chan int, 5)

	//Ініціалізація WaitGroup з лічильником 3 для паралельної обробки
	var wg sync.WaitGroup
	wg.Add(3)

	//Запуск горутини
	for i := 1; i <= 3; i++ {
		go sensorWorker(i, dataCh, &wg)
	}

	//Блокування усіх потоків до завершення роботи горутин
	go func() {
		wg.Wait()
		//Закриття каналу
		close(dataCh)
	}()

	fmt.Println("[Agregator] Waiting for data ...")

	//Вивід результатів
	for value := range dataCh {
		fmt.Printf("[Agregator] reciewed data %d\n", value)
	}

	fmt.Println("Work finished! Channel closed!")
}
