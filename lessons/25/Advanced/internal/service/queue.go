package service

import "queueapp/internal/storage"

type QueueService struct {
	store *storage.MemoryStore
}

// Функція доступу до даних модулю
func NewQueueService(store *storage.MemoryStore) *QueueService {
	return &QueueService{store: store}
}

//Функція створення запису клієнта (надає доступ іншим модулям)
func (s *QueueService) AddClient(name string) {
	s.store.Create(name)
}

//Функція отриання усіх записів клієнтів (надає доступ іншим модулям)
func (s *QueueService) GetAll() []storage.QueueItem {
	return s.store.ReadAll()
}

//Функція редагування даних клієнта (надає доступ іншим модулям)
func (s *QueueService) EditClient(id int, name string) error {
	return s.store.Update(id, name)
}

//Функція видалення запису клієнта (надає доступ іншим модулям)
func (s *QueueService) RemoveClient(id int) error {
	return s.store.Delete(id)
}
