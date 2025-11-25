package queue

import (
	"e-queue/internal/logger"
	"fmt"
)

// Структура черги
type Queue struct {
	items []string
	limit int
}

// Функція New утворює чергу
func New(limit int) *Queue {
	logger.Log.Info("Initialized new queue!")
	return &Queue{limit: limit}
}

// Функція Enqueue додає елементи до черги, якщо її максимальна вмісткість ще не досягнута
func (q *Queue) Enqueue(item string) {
	if len(q.items) >= q.limit {
		logger.Log.Warn("Queue is full!")
		return
	}
	q.items = append(q.items, item)
	logger.Log.Infof("Added: %s", item)
}

// Фуекція Dequeue видаляє перший елемент черги
func (q *Queue) Dequeue() (string, bool) {
	if len(q.items) == 0 {
		return "", false
	}
	item := q.items[0]
	q.items = q.items[1:]
	logger.Log.Infof("Deleted: %s", item)
	return item, true
}

// Функція Size виводить кількість елементів черги
func (q *Queue) Size() int {
	return len(q.items)
}

// Функція String виводить вміст черги
func (q *Queue) String() string {
	return fmt.Sprintf("Queue(%v)", q.items)
}
