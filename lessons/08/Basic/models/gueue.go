package models

type QueueItem struct {
	ID       int
	Client   string
	Serviced bool
}

var queueCounter int

func NewQueueItem(client string) QueueItem {
	queueCounter++
	return QueueItem{
		ID:       queueCounter,
		Client:   client,
		Serviced: false,
	}
}
