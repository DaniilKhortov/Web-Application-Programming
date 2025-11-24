package main

import (
	"fmt"
	"time"
)

func sensorWorker(dataCh chan float64) {
	power := 7.5
	fmt.Println("Sending data (power)...")

	dataCh <- power

	fmt.Println("Data reciewed (power)")
}

func main() {

	dataCh := make(chan float64)

	go sensorWorker(dataCh)

	fmt.Println("[Agregator] Waiting for data (2 seconds)...")
	time.Sleep(2 * time.Second)

	powerValue := <-dataCh
	fmt.Printf("[Agregator] reciewed power %.2f \n", powerValue)

	fmt.Println("Work finished!")
}
