package main

import (
	"fmt"
	"sync"
	"time"
)

// Структура стану черги
type QueueState struct {
	sync.RWMutex
	counter int
}

// Функція зчитування стану черги
func (q *QueueState) Get() int {
	q.RLock()
	defer q.RUnlock()

	// імітація читання
	time.Sleep(10 * time.Millisecond)
	return q.counter
}

// Функція оновлення стану
func (q *QueueState) Update() {
	q.Lock()
	defer q.Unlock()

	time.Sleep(50 * time.Millisecond) // імітація запису
	q.counter++
}

func main() {
	var (
		queue QueueState
		wg    sync.WaitGroup
	)

	readers := 100
	writers := 5

	// Запуск readers
	wg.Add(readers)
	for i := 0; i < readers; i++ {
		go func(id int) {
			defer wg.Done()
			value := queue.Get()
			fmt.Printf("Reader %d read value: %d\n", id, value)
		}(i)
	}

	// Запуск writers
	wg.Add(writers)
	for i := 0; i < writers; i++ {
		go func(id int) {
			defer wg.Done()
			queue.Update()
			fmt.Printf("Writer %d updated queue\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("Final queue state:", queue.counter)
}
