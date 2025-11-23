package storage

import (
	"errors"
	"math"
)

type QueueItem struct {
	ID   int
	Name string
}

type MemoryStore struct {
	data []QueueItem
	next int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{data: []QueueItem{}, next: 1}
}

func (m *MemoryStore) Create(name string) {
	id := int(math.Abs(float64(m.next)))
	m.data = append(m.data, QueueItem{ID: id, Name: name})
	m.next++
}

func (m *MemoryStore) ReadAll() []QueueItem {
	return m.data
}

func (m *MemoryStore) Update(id int, newName string) error {
	for i, item := range m.data {
		if item.ID == id {
			m.data[i].Name = newName
			return nil
		}
	}
	return errors.New("item not found")
}

func (m *MemoryStore) Delete(id int) error {
	for i, item := range m.data {
		if item.ID == id {
			m.data = append(m.data[:i], m.data[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}
