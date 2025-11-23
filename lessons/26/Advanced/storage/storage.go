package storage

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

// Client — модель клієнта черги
type Client struct {
	ID   int
	Name string
}

// Storage — структура для взаємодії з БД
type Storage struct {
	DB *sql.DB
}

// NewStorage створює підключення до БД (sqlite для простоти)
func NewStorage(dbPath string) (*Storage, error) {
	db, err := sql.Open("sqlite", dbPath) // ← ось тут зміна!
	if err != nil {
		return nil, fmt.Errorf("cannot open DB: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("cannot connect to DB: %w", err)
	}

	return &Storage{DB: db}, nil
}

// Init створює таблицю, якщо її ще немає
func (s *Storage) Init() error {
	query := `
	CREATE TABLE IF NOT EXISTS clients (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	);`
	_, err := s.DB.Exec(query)
	return err
}

// Clear очищає таблицю (для тестів)
func (s *Storage) Clear() error {
	_, err := s.DB.Exec("DELETE FROM clients;")
	return err
}

// AddClient додає нового клієнта в чергу
func (s *Storage) AddClient(name string) (int64, error) {
	res, err := s.DB.Exec("INSERT INTO clients(name) VALUES(?);", name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// GetClientByID отримує клієнта за ID
func (s *Storage) GetClientByID(id int64) (*Client, error) {
	row := s.DB.QueryRow("SELECT id, name FROM clients WHERE id = ?;", id)
	var c Client
	if err := row.Scan(&c.ID, &c.Name); err != nil {
		return nil, err
	}
	return &c, nil
}
