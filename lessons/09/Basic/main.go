package main

import (
	"fmt"
)

// Інтерфейс функцій роботи з чергою
type Queueable interface {
	AddClient(name string)
	NextClient() string
	Count() int
}

// Структура звичайної черги
type SimpleQueue struct {
	clients []string
}

// Функція додавання клієнтів до черги
func (q *SimpleQueue) AddClient(name string) {
	q.clients = append(q.clients, name)
	fmt.Printf("Added clients %s into SimpleQueue\n", name)
}

// Функція обслуговування клієнтів до черги
func (q *SimpleQueue) NextClient() string {
	if len(q.clients) == 0 {
		return "Queue is empty"
	}
	next := q.clients[0]
	q.clients = q.clients[1:]
	return next
}

// Функція обчислення розміру пріоритетної черги
func (q *SimpleQueue) Count() int {
	return len(q.clients)
}

// Структура пріоритетної черги
type PriorityQueue struct {
	priorityClients []string
	normalClients   []string
}

// Функція додавання клієнта до пріоритетної черги
func (pq *PriorityQueue) AddClient(name string) {
	if len(name) > 0 && name[0] == '!' {
		pq.priorityClients = append(pq.priorityClients, name[1:])
		fmt.Printf("High priority Client %s added\n", name[1:])
	} else {
		pq.normalClients = append(pq.normalClients, name)
		fmt.Printf("Client %s added into PriorityQueue\n", name)
	}
}

// Функція обслуговування клієнтів до пріоритетної черги
func (pq *PriorityQueue) NextClient() string {
	if len(pq.priorityClients) > 0 {
		next := pq.priorityClients[0]
		pq.priorityClients = pq.priorityClients[1:]
		return next + " (high priority)"
	}
	if len(pq.normalClients) > 0 {
		next := pq.normalClients[0]
		pq.normalClients = pq.normalClients[1:]
		return next
	}
	return "Queue is empty"
}

// Функція обчислення розміру пріоритетної черги
func (pq *PriorityQueue) Count() int {
	return len(pq.priorityClients) + len(pq.normalClients)
}

// Функція виводу розміру пріоритетної черги
func (pq PriorityQueue) String() string {
	return fmt.Sprintf("PriorityQueue: %d high priority, %d common clients",
		len(pq.priorityClients), len(pq.normalClients))
}

// Функція обслуговування клієнтів черги
func ServeNext(q Queueable) {
	fmt.Println("Servicing:", q.NextClient())

	switch v := q.(type) {
	case *SimpleQueue:
		fmt.Printf("Clients in SimpleQueue left: %d\n", v.Count())
	case *PriorityQueue:
		fmt.Printf("Clients in PriorityQueue left: %d\n", v.Count())
	default:
		fmt.Println("Unknown Queue type")
	}
}

func main() {
	//Утворення черг
	sq := &SimpleQueue{}
	pq := &PriorityQueue{}

	//Додавання клієнтів до черг
	sq.AddClient("Oleh")
	sq.AddClient("Maksym")
	pq.AddClient("!Andriy")
	pq.AddClient("Sosza")

	fmt.Println()

	//Обслуговування черг
	ServeNext(sq)
	ServeNext(pq)

	fmt.Println()
	fmt.Println("PriorityQueue info:", pq)
}
