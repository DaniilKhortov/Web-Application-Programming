package main

import (
	"fmt"
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

type FullQueue interface {
	Queueable
	Monitorable
	fmt.Stringer
}

type BaseQueue struct {
	clients []string
}

func (b *BaseQueue) AddClient(name string) {
	b.clients = append(b.clients, name)
}

func (b *BaseQueue) NextClient() string {
	if len(b.clients) == 0 {
		return "Queue is empty"
	}
	next := b.clients[0]
	b.clients = b.clients[1:]
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
	} else {
		p.clients = append(p.clients, name)
	}
}

func (p *PriorityQueue) NextClient() string {
	if len(p.priorityClients) > 0 {
		next := p.priorityClients[0]
		p.priorityClients = p.priorityClients[1:]
		return next + " (high priority)"
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
	sq := &SimpleQueue{}
	pq := &PriorityQueue{}

	sq.AddClient("Oleh")
	sq.AddClient("Maksym")
	pq.AddClient("!Andriy")
	pq.AddClient("Sosza")

	queues := []FullQueue{sq, pq}

	ServeAll(queues)
}
