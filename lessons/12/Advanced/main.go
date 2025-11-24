package main

import (
	"fmt"
	"time"
)

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

func consumer(dataCh <-chan int) {
	for value := range dataCh {
		fmt.Printf("Consumer: recieved data %d\n", value)
		time.Sleep(700 * time.Millisecond)
	}
	fmt.Println("Consumer: channel is closed! Work finished!")
}

func main() {

	dataCh := make(chan int, 3)

	done := make(chan bool)

	go producer(dataCh, done)
	go consumer(dataCh)

	time.Sleep(3 * time.Second)
	fmt.Println("Main: sending sygnal of finishing to producer...")
	done <- true

	time.Sleep(1 * time.Second)
	close(dataCh)

	fmt.Println("Main: work finished!")
}
