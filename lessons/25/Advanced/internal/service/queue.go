package service

import "queueapp/internal/storage"

type QueueService struct {
	store *storage.MemoryStore
}

func NewQueueService(store *storage.MemoryStore) *QueueService {
	return &QueueService{store: store}
}

func (s *QueueService) AddClient(name string) {
	s.store.Create(name)
}

func (s *QueueService) GetAll() []storage.QueueItem {
	return s.store.ReadAll()
}

func (s *QueueService) EditClient(id int, name string) error {
	return s.store.Update(id, name)
}

func (s *QueueService) RemoveClient(id int) error {
	return s.store.Delete(id)
}
