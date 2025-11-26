package main

import (
	"fmt"
	"time"
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

// Інтерфейс функцій відлагодження
type Loggable interface {
	LogAction(action string)
	GetLog() []string
}

// Структура звичайної черги
type BaseQueue struct {
	clients []string
	logs    []string
}

// Інтерфейс AdvancedQueue є композицією всіх інших інтерфесів
// Такий підхід полегшує доступ до функцій різних інтерфейсів
type AdvancedQueue interface {
	Queueable
	Monitorable
	Loggable
	fmt.Stringer
}

// Функція додавання клієнтів до черги
func (b *BaseQueue) AddClient(name string) {
	b.clients = append(b.clients, name)
	b.LogAction(fmt.Sprintf("Added client %s", name))
}

// Функція обслуговування клієнтів до черги
func (b *BaseQueue) NextClient() string {
	if len(b.clients) == 0 {
		b.LogAction("Queue is empty")
		return "Queue is empty"
	}
	next := b.clients[0]
	b.clients = b.clients[1:]
	b.LogAction(fmt.Sprintf("Serviced client %s", next))
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

// Функція отримання даних усіх клієнтів з черги
func (b *BaseQueue) GetAll() []string {
	return b.clients
}

// Функція задокументовування виконаних дій
func (b *BaseQueue) LogAction(action string) {
	entry := fmt.Sprintf("[%s] %s", time.Now().Format("15:04:05"), action)
	b.logs = append(b.logs, entry)
}

// Функція отримання даних логера
func (b *BaseQueue) GetLog() []string {
	return b.logs
}

// Структура звичайної черги
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
		p.LogAction(fmt.Sprintf("Added high priority client:  %s", name[1:]))
	} else {
		p.clients = append(p.clients, name)
		p.LogAction(fmt.Sprintf("Added common client: %s", name))
	}
}

// Функція обслуговування клієнтів до пріоритетної черги
func (p *PriorityQueue) NextClient() string {
	if len(p.priorityClients) > 0 {
		next := p.priorityClients[0]
		p.priorityClients = p.priorityClients[1:]
		p.LogAction(fmt.Sprintf("Serviced high priority client :%s", next))
		return next + " (пріоритетний)"
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

// Універсальна функція виводу характеристик черги
func PrintAnyQueueInfo(obj interface{}) {
	switch v := obj.(type) {
	case AdvancedQueue:
		fmt.Println("Current queue:", v)
		fmt.Println("Next:", v.Peek())
		fmt.Println("History log:")
		for _, log := range v.GetLog() {
			fmt.Println("   ", log)
		}
	case string:
		fmt.Println("Order:", v)
	default:
		fmt.Printf("Unknown type (%T)\n", v)
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

	//Ініціалізація інтерфейсу AdvancedQueue для обробки черг
	var queues []AdvancedQueue = []AdvancedQueue{sq, pq}

	//Обслуговування черг
	for _, q := range queues {
		q.NextClient()
	}

	//Виведення результатів
	fmt.Println("\nEmpty interface output:")
	PrintAnyQueueInfo(sq)
	fmt.Println()
	PrintAnyQueueInfo(pq)
	fmt.Println()
	PrintAnyQueueInfo("PrintAnyQueueInfo test")
}
