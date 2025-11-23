package main

import (
	"fmt"
	"time"
)

// sensorWorker імітує датчик потужності в системі електронної черги
func sensorWorker(dataCh chan float64) {
	power := 7.5
	fmt.Println("Sending data (power)...")

	// Надсилаємо дані у канал (блокується, якщо немає приймача)
	dataCh <- power

	fmt.Println("Data reciewed (power)")
}

func main() {
	// Створення небуферизованого каналу для передачі float64
	dataCh := make(chan float64)

	// Запуск goroutine-датчика
	go sensorWorker(dataCh)

	// Імітація затримки — приймач ще "не готовий"
	fmt.Println("[Agregator] Waiting for data (2 seconds)...")
	time.Sleep(2 * time.Second)

	// Приймання даних з каналу
	powerValue := <-dataCh
	fmt.Printf("[Agregator] reciewed power %.2f \n", powerValue)

	fmt.Println("Work finished!")
}
