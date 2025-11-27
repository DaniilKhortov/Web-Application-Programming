package main

import (
	"fmt"
	"time"
)

// Горутина producer. Вона виводитиме дані до отримання сигналу завершення роботи
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

// Горутина consumer. Вона виводитиме дані буферизованого каналу до його закриття
func consumer(dataCh <-chan int) {
	for value := range dataCh {
		fmt.Printf("Consumer: recieved data %d\n", value)
		time.Sleep(700 * time.Millisecond)
	}
	fmt.Println("Consumer: channel is closed! Work finished!")
}

func main() {

	//Ініціалізація буфериованого каналу
	dataCh := make(chan int, 3)

	//Ініціалізація звичайного каналу
	done := make(chan bool)

	//Запуск горутин
	go producer(dataCh, done)
	go consumer(dataCh)

	time.Sleep(3 * time.Second)
	fmt.Println("Main: sending sygnal of finishing to producer...")

	//Після завершення роботи, передаємо сигнал завершення роботи
	done <- true

	time.Sleep(1 * time.Second)

	//Закриття каналу
	close(dataCh)

	fmt.Println("Main: work finished!")
}
