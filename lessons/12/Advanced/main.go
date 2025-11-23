package main

import (
	"fmt"
	"time"
)

// producer — функція, що відправляє дані у канал лише для запису
// та реагує на сигнал завершення через канал лише для читання
func producer(dataCh chan<- int, done <-chan bool) {
	for i := 1; i <= 10; i++ {
		select {
		case <-done:
			fmt.Println("Producer: sygnal of finishing recieved. Work finished!")
			return
		case dataCh <- i:
			fmt.Printf("Producer: sent data %d\n", i)
			time.Sleep(400 * time.Millisecond)
		}
	}
	fmt.Println("Producer: all data sent. That is all for today!")
}

// consumer — функція, що приймає дані з каналу лише для читання
func consumer(dataCh <-chan int) {
	for value := range dataCh {
		fmt.Printf("Consumer: recieved data %d\n", value)
		time.Sleep(700 * time.Millisecond)
	}
	fmt.Println("Consumer: channel is closed! Work finished!")
}

func main() {
	// Канал для даних (буферизований, щоб уникнути блокувань при коротких затримках)
	dataCh := make(chan int, 3)

	// Канал сигналу завершення
	done := make(chan bool)

	// Запуск producer та consumer
	go producer(dataCh, done)
	go consumer(dataCh)

	// Даємо системі попрацювати 3 секунди, а потім зупиняємо producer
	time.Sleep(3 * time.Second)
	fmt.Println("Main: sending sygnal of finishing to producer...")
	done <- true

	// Трохи чекаємо, щоб завершились всі операції
	time.Sleep(1 * time.Second)
	close(dataCh)

	fmt.Println("Main: work finished!")
}
