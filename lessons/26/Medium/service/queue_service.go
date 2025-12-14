package service

// QueueService визначає інтерфейс для роботи з чергою клієнтів.
type QueueService interface {
	GetAllClients() []string
}

// Ініціалізація RealQueueService
type RealQueueService struct{}

//Функція повернення даних клієнтів
func (r *RealQueueService) GetAllClients() []string {

	return []string{"Olha", "Ivan", "Maria"}
}
