package main

import (
	"fmt"
	"time"
)

// Горутина, що надсилає дані у канал
func sensorWorker(dataCh chan float64) {
	power := 7.5
	fmt.Println("Sending data (power)...")

	dataCh <- power

	fmt.Println("Data reciewed (power)")
}

func main() {

	//Створення каналу даних
	dataCh := make(chan float64)

	//Запуск горутини
	go sensorWorker(dataCh)

	//Імітація очікування на результат
	fmt.Println("[Agregator] Waiting for data (2 seconds)...")
	time.Sleep(2 * time.Second)

	//Збереження даних каналу у змінну
	powerValue := <-dataCh
	fmt.Printf("[Agregator] reciewed power %.2f \n", powerValue)

	fmt.Println("Work finished!")
}
