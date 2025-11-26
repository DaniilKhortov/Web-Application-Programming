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

// Інтерфейс функцій виводу черги
type Monitorable interface {
	Peek() string
	GetAll() []string
}

// Інтерфейс FullQueue є композицією всіх інших інтерфесів
// Такий підхід полегшує доступ до функцій різних інтерфейсів
type FullQueue interface {
	Queueable
	Monitorable
	fmt.Stringer
}

// Структура звичайної черги
type BaseQueue struct {
	clients []string
}

// Функція додавання клієнтів до черги
func (b *BaseQueue) AddClient(name string) {
	b.clients = append(b.clients, name)
}

// Функція обслуговування клієнтів до черги
func (b *BaseQueue) NextClient() string {
	if len(b.clients) == 0 {
		return "Queue is empty"
	}
	next := b.clients[0]
	b.clients = b.clients[1:]
	return next
}

// Функція обчислення розміру черги
func (b *BaseQueue) Count() int {
	return len(b.clients)
}

// Функція перегляду першого клієнта в черзі
func (b *BaseQueue) Peek() string {
	if len(b.clients) == 0 {
		return "No clients left!"
	}
	return b.clients[0]
}

func (b *BaseQueue) GetAll() []string {
	return b.clients
}

// Функція отримання даних усіх клієнтів з черги
type SimpleQueue struct {
	BaseQueue
}

// Функція виводу розміру звичайної черги
func (s SimpleQueue) String() string {
	return fmt.Sprintf("SimpleQueue (%d clients)", s.Count())
}

// Структура пріоритетної черги
type PriorityQueue struct {
	BaseQueue
	priorityClients []string
}

// Функція додавання клієнта до пріоритетної черги
func (p *PriorityQueue) AddClient(name string) {
	if len(name) > 0 && name[0] == '!' {
		p.priorityClients = append(p.priorityClients, name[1:])
	} else {
		p.clients = append(p.clients, name)
	}
}

// Функція обслуговування клієнтів до пріоритетної черги
func (p *PriorityQueue) NextClient() string {
	if len(p.priorityClients) > 0 {
		next := p.priorityClients[0]
		p.priorityClients = p.priorityClients[1:]
		return next + " (high priority)"
	}
	return p.BaseQueue.NextClient()
}

// Функція обчислення розміру пріоритетної черги
func (p *PriorityQueue) Count() int {
	return len(p.priorityClients) + len(p.clients)
}

// Функція перегляду першого клієнта в пріоритетній черзі
func (p *PriorityQueue) Peek() string {
	if len(p.priorityClients) > 0 {
		return p.priorityClients[0] + " (high priority)"
	}
	return p.BaseQueue.Peek()
}

// Функція виводу розміру пріоритетної черги
func (p PriorityQueue) String() string {
	return fmt.Sprintf("PriorityQueue: %d high priority, %d common clients",
		len(p.priorityClients), len(p.clients))
}

// Функція обслуговування усіх клієнтів
func ServeAll(queues []FullQueue) {
	for _, q := range queues {
		fmt.Println("Current queue:", q)
		fmt.Println("Servicing:", q.NextClient())

		switch v := q.(type) {
		case *SimpleQueue:
			fmt.Println("Clients in SimpleQueue left:", v.Count())
		case *PriorityQueue:
			fmt.Println("Clients in PriorityQueue left:", v.Count())
		default:
			fmt.Println("Unknown Queue type")
		}

		fmt.Println("Next:", q.Peek())
		fmt.Println()
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

	//Ініціалізація інтерфейсу FullQueue для обробки черг
	queues := []FullQueue{sq, pq}

	//Обслуговування черг
	ServeAll(queues)
}
