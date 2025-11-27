package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Структура збереження результатів
type Result struct {
	id     int
	status string
}

// Горутина обслуговування клієнта
func serveClient(id int, results chan<- Result) {
	start := time.Now()
	fmt.Printf("Goroutine %d starting (client in queue)\n", id)

	//Імітація обробки запиту клієнта
	time.Sleep(time.Duration(100+rand.Intn(200)) * time.Millisecond)

	//Заповнення результатів через канал
	results <- Result{
		id:     id,
		status: fmt.Sprintf("Client %d was served by %v", id, time.Since(start)),
	}
	fmt.Printf("Goroutine %d finished (client served)\n", id)
}

// Функція обслуговування клієнтів у кількох потоках
func parallelProcessing(numClients int) time.Duration {
	start := time.Now()

	//Ініціалізація мапи результатів
	results := make(chan Result, numClients)
	defer close(results)

	//Запуск горутини
	for i := 1; i <= numClients; i++ {
		go serveClient(i, results)
	}

	//Відслідковування виконаних горутин
	done := 0
	for done < numClients {
		select {
		case res := <-results:
			fmt.Println("-->", res.status)
			done++
		}
	}

	//Вивід результатів
	elapsed := time.Since(start)
	fmt.Printf("\n[Paralel] All clients served for %v\n", elapsed)
	return elapsed
}

// Функція обслуговування клієнтів в одному потоці
func sequentialProcessing(numClients int) time.Duration {
	start := time.Now()

	//Обслуговування клієнтів
	for i := 1; i <= numClients; i++ {
		fmt.Printf("Servicing client %d gradually ...\n", i)
		time.Sleep(time.Duration(100+rand.Intn(200)) * time.Millisecond)
	}

	//Вивід результатів
	elapsed := time.Since(start)
	fmt.Printf("\n[Gradual] All clients served for %v\n", elapsed)
	return elapsed
}

func main() {

	//Задання кількості клієнтів
	numClients := 8

	//Обслуговуваггя клієнтів через конкурентне виконання
	fmt.Println("Paralel computing")
	parallelTime := parallelProcessing(numClients)

	//Обслуговуваггя клієнтів через послідовне виконання
	fmt.Println("Gradual computing")
	sequentialTime := sequentialProcessing(numClients)

	//Вивід результатів
	fmt.Println("\nResults")
	fmt.Printf("Paralel computing: %v\n", parallelTime)
	fmt.Printf("Gradual computing: %v\n", sequentialTime)
	fmt.Printf("Effectivness:  %.2fx\n", float64(sequentialTime)/float64(parallelTime))
}
