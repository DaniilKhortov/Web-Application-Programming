package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	id     int
	status string
}

func serveClient(id int, results chan<- Result) {
	start := time.Now()
	fmt.Printf("Goroutine %d starting (client in queue)\n", id)

	time.Sleep(time.Duration(100+rand.Intn(200)) * time.Millisecond)

	results <- Result{
		id:     id,
		status: fmt.Sprintf("Client %d was served by %v", id, time.Since(start)),
	}
	fmt.Printf("Goroutine %d finished (client served)\n", id)
}

func parallelProcessing(numClients int) time.Duration {
	start := time.Now()

	results := make(chan Result, numClients)
	defer close(results)

	for i := 1; i <= numClients; i++ {
		go serveClient(i, results)
	}

	done := 0
	for done < numClients {
		select {
		case res := <-results:
			fmt.Println("-->", res.status)
			done++
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("\n[Paralel] All clients served for %v\n", elapsed)
	return elapsed
}

func sequentialProcessing(numClients int) time.Duration {
	start := time.Now()

	for i := 1; i <= numClients; i++ {
		fmt.Printf("Servicing client %d gradually ...\n", i)
		time.Sleep(time.Duration(100+rand.Intn(200)) * time.Millisecond)
	}

	elapsed := time.Since(start)
	fmt.Printf("\n[Gradual] All clients served for %v\n", elapsed)
	return elapsed
}

func main() {

	numClients := 8

	fmt.Println("Paralel computing")

	parallelTime := parallelProcessing(numClients)

	fmt.Println("Gradual computing")
	sequentialTime := sequentialProcessing(numClients)

	fmt.Println("\nResults")
	fmt.Printf("Paralel computing: %v\n", parallelTime)
	fmt.Printf("Gradual computing: %v\n", sequentialTime)
	fmt.Printf("Effectivness:  %.2fx\n", float64(sequentialTime)/float64(parallelTime))
}
