package interfaces

//Інтерфейс QueueService. Використовуєтья для доступу до методів іншого модуля
type QueueService interface {
	Enqueue(item string)
	Dequeue() (string, bool)
	Size() int
}
