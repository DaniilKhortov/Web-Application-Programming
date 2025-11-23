// file: main.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// processingDone — канал повідомляє, що черговий клієнт оброблений
// shutdown — глобальний сигнал зупинки всієї системи (вимикає монітор)
func main() {

	processingDone := make(chan int) // небуферизований канал: події завершення обробки
	shutdown := make(chan struct{})  // небуферизований канал: глобальне завершення

	// Імітація джерела подій (обробка клієнтів)
	go clientProcessor(processingDone, shutdown)

	// Імітація монітору стану системи
	go statusMonitor(processingDone, shutdown)

	// Дозволяємо системі попрацювати кілька секунд
	time.Sleep(3 * time.Second)

	// Відправляємо глобальний сигнал завершення
	fmt.Println("\n[MAIN] Sending command to end...")
	close(shutdown)

	// Додатковий час для коректного завершення горутин
	time.Sleep(1 * time.Second)
	fmt.Println("[MAIN] Work is done!")
}

// clientProcessor — генерує події (обробку клієнтів)
func clientProcessor(done chan<- int, shutdown <-chan struct{}) {
	clientID := 1
	for {
		select {
		case <-shutdown:
			fmt.Println("[Processor] Recieved finnishing command. Aborting process.")
			return
		default:
			// імітація тривалості обробки клієнта
			time.Sleep(time.Duration(200+rand.Intn(300)) * time.Millisecond)
			fmt.Printf("[Processor] Client #%d served.\n", clientID)
			done <- clientID
			clientID++
		}
	}
}

// statusMonitor — слухає два канали: події обробки і сигнал завершення
func statusMonitor(done <-chan int, shutdown <-chan struct{}) {
	active := 0
	for {
		select {
		case id := <-done:
			active++
			fmt.Printf("[Monitor] Clients served: %d (last #%d)\n", active, id)

		case <-shutdown:
			fmt.Printf("[Monitor] Recieved finnishing command. Result: %d clients.\n", active)
			return
		}
	}
}
