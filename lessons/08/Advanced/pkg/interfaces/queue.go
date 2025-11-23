package interfaces

// QueueService визначає загальний інтерфейс для систем черг.
type QueueService interface {
	Enqueue(item string)
	Dequeue() (string, bool)
	Size() int
}
