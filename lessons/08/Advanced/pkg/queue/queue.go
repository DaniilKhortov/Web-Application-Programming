package queue

import (
	"e-queue/internal/logger"
	"fmt"
)

// Queue реалізує просту чергу рядків.
type Queue struct {
	items []string
	limit int
}

// New створює нову чергу з вказаним розміром.
func New(limit int) *Queue {
	logger.Log.Info("Initialized new queue!")
	return &Queue{limit: limit}
}

// Enqueue додає елемент у чергу.
func (q *Queue) Enqueue(item string) {
	if len(q.items) >= q.limit {
		logger.Log.Warn("Queue is full!")
		return
	}
	q.items = append(q.items, item)
	logger.Log.Infof("Added: %s", item)
}

// Dequeue витягує елемент із черги.
func (q *Queue) Dequeue() (string, bool) {
	if len(q.items) == 0 {
		return "", false
	}
	item := q.items[0]
	q.items = q.items[1:]
	logger.Log.Infof("Deleted: %s", item)
	return item, true
}

// Size повертає поточну довжину черги.
func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) String() string {
	return fmt.Sprintf("Queue(%v)", q.items)
}
