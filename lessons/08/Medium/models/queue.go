package models

// QueueItem представляє одного клієнта в електронній черзі.
// Це експортована структура (видима з інших пакетів).
type QueueItem struct {
	ID       int
	Client   string
	Serviced bool
}
