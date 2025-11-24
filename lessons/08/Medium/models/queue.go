package models

type QueueItem struct {
	ID       int
	Client   string
	Serviced bool
}
