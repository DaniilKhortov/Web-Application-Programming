package models

// QueueItem — експортована структура, видима з інших пакетів
type QueueItem struct {
	ID       int
	Client   string
	Serviced bool
}

// queueCounter — приватний лічильник, не експортується
var queueCounter int

// NewQueueItem — створює новий елемент черги
func NewQueueItem(client string) QueueItem {
	queueCounter++
	return QueueItem{
		ID:       queueCounter,
		Client:   client,
		Serviced: false,
	}
}
