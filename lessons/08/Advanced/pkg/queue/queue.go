package queue

import (
	"e-queue/internal/logger"
	"fmt"
)

type Queue struct {
	items []string
	limit int
}

func New(limit int) *Queue {
	logger.Log.Info("Initialized new queue!")
	return &Queue{limit: limit}
}

func (q *Queue) Enqueue(item string) {
	if len(q.items) >= q.limit {
		logger.Log.Warn("Queue is full!")
		return
	}
	q.items = append(q.items, item)
	logger.Log.Infof("Added: %s", item)
}

func (q *Queue) Dequeue() (string, bool) {
	if len(q.items) == 0 {
		return "", false
	}
	item := q.items[0]
	q.items = q.items[1:]
	logger.Log.Infof("Deleted: %s", item)
	return item, true
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) String() string {
	return fmt.Sprintf("Queue(%v)", q.items)
}
