package service

// QueueService визначає інтерфейс для роботи з чергою клієнтів.
type QueueService interface {
	GetAllClients() []string
}

// RealQueueService — справжня реалізація (для реального проєкту може брати дані з БД)
type RealQueueService struct{}

func (r *RealQueueService) GetAllClients() []string {
	// У реальному застосунку тут був би запит до БД.
	return []string{"Olha", "Ivan", "Maria"}
}
