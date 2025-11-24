package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type Address struct {
	City    string `json:"city"`
	Street  string `json:"street"`
	ZipCode string `json:"zip_code"`
}

func NewAddress(city, street, zip string) Address {
	return Address{City: city, Street: street, ZipCode: zip}
}

type QueueData struct {
	Position int `json:"position"`
	AvgTime  int `json:"avg_time"`
}

func NewQueueData(pos, avg int) QueueData {
	return QueueData{Position: pos, AvgTime: avg}
}

func (q QueueData) EstimatedWait() int {
	return q.Position * q.AvgTime
}

func (q *QueueData) SetPosition(p int) { q.Position = p }
func (q *QueueData) SetAvgTime(a int)  { q.AvgTime = a }

type Client struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Phone   string  `json:"phone"`
	Address Address `json:"address"`
	QueueData
}

func NewClient(id int, name, email, phone string, addr Address, pos, avg int) Client {
	return Client{
		ID:        id,
		Name:      name,
		Email:     email,
		Phone:     phone,
		Address:   addr,
		QueueData: NewQueueData(pos, avg),
	}
}

func (c Client) Info() string {
	return fmt.Sprintf(
		"Client #%d: %s\nEmail: %s | Phone: %s\nAddress: %s, %s (%s)\nPosition: %d | Wait: %d min\n",
		c.ID, c.Name, c.Email, c.Phone,
		c.Address.City, c.Address.Street, c.Address.ZipCode,
		c.Position, c.EstimatedWait(),
	)
}

func (c *Client) UpdatePosition(newPos int) {
	c.Position = newPos
}

type QueueManager struct {
	clients []Client
}

func NewQueueManager() *QueueManager {
	return &QueueManager{clients: make([]Client, 0)}
}

func (qm *QueueManager) AddClient(c Client) {
	qm.clients = append(qm.clients, c)
}

func (qm *QueueManager) TotalEstimatedTime() int {
	total := 0
	for _, c := range qm.clients {
		total += c.EstimatedWait()
	}
	return total
}

func (qm *QueueManager) Reorder() {
	sort.SliceStable(qm.clients, func(i, j int) bool {
		return qm.clients[i].Position < qm.clients[j].Position
	})
}

func (qm QueueManager) ToJSON() string {
	data, _ := json.MarshalIndent(qm.clients, "", "  ")
	return string(data)
}

func (qm QueueManager) PrintAll() {
	for _, c := range qm.clients {
		fmt.Println(c.Info())
	}
}

func main() {
	fmt.Printf("E-Queue\n\n")

	addr1 := NewAddress("Kyiv", "Khreshchatyk", "01001")
	addr2 := NewAddress("Lviv", "Svobody Ave", "79000")

	client1 := NewClient(1, "Ivan Ivanov", "ivan@example.com", "+380501234567", addr1, 3, 4)
	client2 := NewClient(2, "Olena Petrova", "olena@example.com", "+380501112233", addr2, 1, 6)

	manager := NewQueueManager()
	manager.AddClient(client1)
	manager.AddClient(client2)

	fmt.Println("Queue:")
	manager.PrintAll()

	manager.Reorder()

	fmt.Println("\nAfter updating positions:")
	manager.PrintAll()

	fmt.Printf("Total estimated waiting time: %d minutes\n", manager.TotalEstimatedTime())

	fmt.Println("\nJSON Export:")
	fmt.Println(manager.ToJSON())
}
