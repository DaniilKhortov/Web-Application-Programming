package main

import (
	"fmt"
	"time"
)

type Queueable interface {
	AddClient(name string)
	NextClient() string
	Count() int
}

type Monitorable interface {
	Peek() string
	GetAll() []string
}

type Loggable interface {
	LogAction(action string)
	GetLog() []string
}

type AdvancedQueue interface {
	Queueable
	Monitorable
	Loggable
	fmt.Stringer
}

type BaseQueue struct {
	clients []string
	logs    []string
}

func (b *BaseQueue) AddClient(name string) {
	b.clients = append(b.clients, name)
	b.LogAction(fmt.Sprintf("Added client %s", name))
}

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

func (b *BaseQueue) Count() int {
	return len(b.clients)
}

func (b *BaseQueue) Peek() string {
	if len(b.clients) == 0 {
		return "No clients left!"
	}
	return b.clients[0]
}

func (b *BaseQueue) GetAll() []string {
	return b.clients
}

func (b *BaseQueue) LogAction(action string) {
	entry := fmt.Sprintf("[%s] %s", time.Now().Format("15:04:05"), action)
	b.logs = append(b.logs, entry)
}

func (b *BaseQueue) GetLog() []string {
	return b.logs
}

type SimpleQueue struct {
	BaseQueue
}

func (s SimpleQueue) String() string {
	return fmt.Sprintf("SimpleQueue (%d clients)", s.Count())
}

type PriorityQueue struct {
	BaseQueue
	priorityClients []string
}

func (p *PriorityQueue) AddClient(name string) {
	if len(name) > 0 && name[0] == '!' {
		p.priorityClients = append(p.priorityClients, name[1:])
		p.LogAction(fmt.Sprintf("Added high priority client:  %s", name[1:]))
	} else {
		p.clients = append(p.clients, name)
		p.LogAction(fmt.Sprintf("Added common client: %s", name))
	}
}

func (p *PriorityQueue) NextClient() string {
	if len(p.priorityClients) > 0 {
		next := p.priorityClients[0]
		p.priorityClients = p.priorityClients[1:]
		p.LogAction(fmt.Sprintf("Serviced high priority client :%s", next))
		return next + " (пріоритетний)"
	}
	return p.BaseQueue.NextClient()
}

func (p *PriorityQueue) Count() int {
	return len(p.priorityClients) + len(p.clients)
}

func (p *PriorityQueue) Peek() string {
	if len(p.priorityClients) > 0 {
		return p.priorityClients[0] + " (high priority)"
	}
	return p.BaseQueue.Peek()
}

func (p PriorityQueue) String() string {
	return fmt.Sprintf("PriorityQueue: %d high priority, %d common clients",
		len(p.priorityClients), len(p.clients))
}

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
	sq := &SimpleQueue{}
	pq := &PriorityQueue{}

	sq.AddClient("Oleh")
	sq.AddClient("Maksym")
	pq.AddClient("!Andriy")
	pq.AddClient("Sosza")

	var queues []AdvancedQueue = []AdvancedQueue{sq, pq}

	for _, q := range queues {
		q.NextClient()
	}

	fmt.Println("\nEmpty interface output:")
	PrintAnyQueueInfo(sq)
	fmt.Println()
	PrintAnyQueueInfo(pq)
	fmt.Println()
	PrintAnyQueueInfo("PrintAnyQueueInfo test")
}
