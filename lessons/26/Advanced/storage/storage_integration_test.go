//go:build integration
// +build integration

package storage

import (
	"os"
	"testing"
)

var testDBPath = "test_queue.db"

// setup — створює тестову базу даних
func setup(t *testing.T) *Storage {
	s, err := NewStorage(testDBPath)
	if err != nil {
		t.Fatalf("cannot create storage: %v", err)
	}
	if err := s.Init(); err != nil {
		t.Fatalf("cannot init table: %v", err)
	}
	if err := s.Clear(); err != nil {
		t.Fatalf("cannot clear data: %v", err)
	}
	return s
}

// teardown — видаляє тестову базу після завершення
func teardown(t *testing.T) {
	if err := os.Remove(testDBPath); err != nil && !os.IsNotExist(err) {
		t.Logf("cannot remove test DB: %v", err)
	}
}

func TestStorage_CreateAndRetrieveClient(t *testing.T) {
	// 1. Налаштування середовища
	s := setup(t)
	defer func() {
		s.DB.Close() // 🔹 Закриваємо БД
		teardown(t)  // 🔹 Потім видаляємо файл
	}()

	// 2. Створюємо клієнта
	expectedName := "Olha"
	id, err := s.AddClient(expectedName)
	if err != nil {
		t.Fatalf("failed to insert client: %v", err)
	}

	// 3. Отримуємо клієнта назад
	client, err := s.GetClientByID(id)
	if err != nil {
		t.Fatalf("failed to retrieve client: %v", err)
	}

	// 4. Перевіряємо цілісність даних
	if client.ID != int(id) {
		t.Errorf("expected ID %d, got %d", id, client.ID)
	}
	if client.Name != expectedName {
		t.Errorf("expected name %q, got %q", expectedName, client.Name)
	}
}
