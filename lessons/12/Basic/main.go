package main

import (
	"fmt"
	"time"
)

func main() {
	// Небуферизований канал для float64
	queue := make(chan float64)

	// Горутина для реєстру клієнтів
	go func() {
		// умовне значення клієнта
		clientValue := 3.14

		fmt.Println("Registrator: attempting to add client to queue...")
		queue <- clientValue
		fmt.Println("Registrator: client was added to queue")
	}()

	// Імітація затримки
	fmt.Println("Operator: is not ready for client service")
	time.Sleep(3 * time.Second)

	// Прийом значення з каналу
	value := <-queue
	fmt.Println("Operator: served client with value:", value)

	time.Sleep(1 * time.Second)
	fmt.Println("That is all for today!")
}
