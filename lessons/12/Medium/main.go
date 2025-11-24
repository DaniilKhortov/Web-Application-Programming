package main

import (
	"fmt"
	"sync"
	"time"
)

func sensorWorker(id int, dataCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 4; i++ {
		value := id*10 + i
		dataCh <- value
		fmt.Printf("Counter #%d: sent value %d\n", id, value)
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Printf("Counter #%d: finished work\n", id)
}

func main() {

	dataCh := make(chan int, 5)

	var wg sync.WaitGroup
	wg.Add(3)

	for i := 1; i <= 3; i++ {
		go sensorWorker(i, dataCh, &wg)
	}

	go func() {
		wg.Wait()
		close(dataCh)
	}()

	fmt.Println("[Agregator] Waiting for data ...")

	for value := range dataCh {
		fmt.Printf("[Agregator] reciewed data %d\n", value)
	}

	fmt.Println("Work finished! Channel closed!")
}
