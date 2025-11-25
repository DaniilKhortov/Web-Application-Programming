package models

//Структура елементу черги
type QueueItem struct {
	ID       int
	Client   string
	Serviced bool
}
