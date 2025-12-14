package main

import (
	"fmt"
	queueapp "main/queueapp"
)

// Запуск тестування
func main() {
	output := queueapp.SimulateQueueConcat(10)
	fmt.Println(output)
}
