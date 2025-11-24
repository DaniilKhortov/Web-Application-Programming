package main

import (
	"fmt"
	"sync"
	"time"
)

type ElectronicQueue struct {
	value int
	mu    sync.RWMutex
}

func (q *ElectronicQueue) Get(id int) int {
	q.mu.RLock()
	defer q.mu.RUnlock()

	fmt.Printf("[Reader %02d] Reading: %d\n", id, q.value)
	time.Sleep(10 * time.Millisecond)
	return q.value
}

func (q *ElectronicQueue) Update(id int) {
	q.mu.Lock()
	defer q.mu.Unlock()

	old := q.value
	q.value++
	fmt.Printf("[Writer %02d] Updating %d into %d\n", id, old, q.value)
	time.Sleep(50 * time.Millisecond)
}

func main() {
	queue := &ElectronicQueue{}
	var wg sync.WaitGroup

	readers := 100
	writers := 5

	fmt.Printf("Running %d readers and %d writers to serve clients simulteniosly.\n\n", readers, writers)

	for i := 0; i < readers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			queue.Get(id)
		}(i + 1)
	}

	for i := 0; i < writers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			queue.Update(id)
		}(i + 1)
	}

	wg.Wait()
	fmt.Printf("\nFinal  queue: %d\n", queue.value)
	fmt.Println("RWMutex synchronization complete!")
}
