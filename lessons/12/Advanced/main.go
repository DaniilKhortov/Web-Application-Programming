package main

import (
	"fmt"
	"time"
)

// producer — відправник даних у канал
func producer(dataCh chan<- int, done <-chan struct{}) {
	defer close(dataCh)

	for i := 1; i <= 10; i++ {
		select {
		case <-done:
			fmt.Println("Registrator: recieved signal. Finishing task...")
			return
		case dataCh <- i:
			fmt.Println("Registrator: added client", i)
			time.Sleep(400 * time.Millisecond)
		}
	}
}

// consumer — отримувач даних з каналу
func consumer(dataCh <-chan int) {
	for value := range dataCh {
		fmt.Println("Operator: servicing client", value)
		time.Sleep(700 * time.Millisecond)
	}
	fmt.Println("Operator: queue is empty")
}

func main() {
	dataCh := make(chan int)

	// Канал скасування
	done := make(chan struct{})

	// Запуск producer та consumer
	go producer(dataCh, done)
	go consumer(dataCh)

	// Імітація зовнішнього скасування
	time.Sleep(3 * time.Second)
	fmt.Println("Main goroutine: sending sygnal of stop")
	close(done)

	time.Sleep(2 * time.Second)
	fmt.Println("That is all for today!")
}
