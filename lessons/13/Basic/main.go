package main

import (
	"fmt"
	"sync"
	"time"
)

const workers = 1000

func simulateWithoutMutex() int {
	var counter int
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func(id int) {

			time.Sleep(time.Microsecond * time.Duration(id%5))

			counter = counter + 1
			wg.Done()
		}(i)
	}

	wg.Wait()
	return counter
}

func simulateWithMutex() int {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func(id int) {

			time.Sleep(time.Microsecond * time.Duration(id%5))

			mu.Lock()
			defer mu.Unlock()
			counter = counter + 1
			wg.Done()
		}(i)
	}

	wg.Wait()
	return counter
}

func main() {

	fmt.Printf("Running %d goroutines to serve clients simulteniosly.\n\n", workers)

	fmt.Println("Data Race without Mutex:")
	noMutexCount := simulateWithoutMutex()
	fmt.Printf("Result: %d (expected %d)\n", noMutexCount, workers)
	if noMutexCount != workers {
		fmt.Println("Data race happened!")
	} else {
		fmt.Println("Somehow data race was avoided. Run again to be sure!")
	}
	fmt.Println()

	fmt.Println("Data Race with Mutex:")
	withMutexCount := simulateWithMutex()
	fmt.Printf("Result: %d (expected %d)\n", withMutexCount, workers)
	if withMutexCount == workers {
		fmt.Println("Data race was avoided!")
	} else {
		fmt.Println("Error! Something went wrong!")
	}
}
