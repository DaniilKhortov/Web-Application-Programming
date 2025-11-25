package models

//Структура елементу черги
type QueueItem struct {
	ID       int
	Client   string
	Serviced bool
}

var queueCounter int

//Функція NewQueueItem створює новий елемнт черги.
//Поле ID визначається автоматично
func NewQueueItem(client string) QueueItem {
	queueCounter++
	return QueueItem{
		ID:       queueCounter,
		Client:   client,
		Serviced: false,
	}
}
