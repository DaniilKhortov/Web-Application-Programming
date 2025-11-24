package interfaces

type QueueService interface {
	Enqueue(item string)
	Dequeue() (string, bool)
	Size() int
}
