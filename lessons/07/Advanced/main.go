package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

// Оголошення структури адреси
// Її поля серіалізуються в JSON
type Address struct {
	City    string `json:"city"`
	Street  string `json:"street"`
	ZipCode string `json:"zip_code"`
}

// Функція NewAddress створює та повертає адресу
func NewAddress(city, street, zip string) Address {
	return Address{City: city, Street: street, ZipCode: zip}
}

// Оголошення структури черги
// Її поля також серіалізуються
type QueueData struct {
	Position int `json:"position"`
	AvgTime  int `json:"avg_time"`
}

// Функція NewQueueData створює та повертає чергу
func NewQueueData(pos, avg int) QueueData {
	return QueueData{Position: pos, AvgTime: avg}
}

// Функція EstimatedWait обчислює час очікування
func (q QueueData) EstimatedWait() int {
	return q.Position * q.AvgTime
}

// Інкапсуляція
// Функція SetPosition - це сеттер до поля Position
func (q *QueueData) SetPosition(p int) { q.Position = p }

// Функція SetAvgTime - це сеттер до поля AvgTime
func (q *QueueData) SetAvgTime(a int) { q.AvgTime = a }

// Оголошення структури клієнта
type Client struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Phone   string  `json:"phone"`
	Address Address `json:"address"`
	//QueueData - це анонімне поле. Вона вбудовує поля та методи QueueData до Client
	//Таким чином утворюється композиція
	QueueData
}

// Функція NewClient створює нового клієнта
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

// Функція Info виводить інформацію клієнта
func (c Client) Info() string {
	return fmt.Sprintf(
		"Client #%d: %s\nEmail: %s | Phone: %s\nAddress: %s, %s (%s)\nPosition: %d | Wait: %d min\n",
		c.ID, c.Name, c.Email, c.Phone,
		c.Address.City, c.Address.Street, c.Address.ZipCode,
		c.Position, c.EstimatedWait(),
	)
}

// Функція UpdatePosition задає нову позицію клієнта
func (c *Client) UpdatePosition(newPos int) {
	c.Position = newPos
}

// Структура QueueManager зберігає клієнтів
type QueueManager struct {
	clients []Client
}

// Функція NewQueueManager ініціалізує QueueManager з пустоим вмістом
func NewQueueManager() *QueueManager {
	return &QueueManager{clients: make([]Client, 0)}
}

// Функція AddClient додає клієнта до черги
func (qm *QueueManager) AddClient(c Client) {
	qm.clients = append(qm.clients, c)
}

// Функція TotalEstimatedTime обчислює загальний час для обслуговування черги
func (qm *QueueManager) TotalEstimatedTime() int {
	total := 0
	for _, c := range qm.clients {
		total += c.EstimatedWait()
	}
	return total
}

// Функція Reorder сортує клієнтів у черзі за місцем
func (qm *QueueManager) Reorder() {
	sort.SliceStable(qm.clients, func(i, j int) bool {
		return qm.clients[i].Position < qm.clients[j].Position
	})
}

// Функція ToJSON серіалізує дані клієнтів
func (qm QueueManager) ToJSON() string {
	data, _ := json.MarshalIndent(qm.clients, "", "  ")
	return string(data)
}

// Функція PrintAll виводить усіх клієнтів в черзі
func (qm QueueManager) PrintAll() {
	for _, c := range qm.clients {
		fmt.Println(c.Info())
	}
}

func main() {
	fmt.Printf("E-Queue\n\n")

	//Створення адрес
	addr1 := NewAddress("Kyiv", "Khreshchatyk", "01001")
	addr2 := NewAddress("Lviv", "Svobody Ave", "79000")

	//Створення клієнтів
	client1 := NewClient(1, "Ivan Ivanov", "ivan@example.com", "+380501234567", addr1, 3, 4)
	client2 := NewClient(2, "Olena Petrova", "olena@example.com", "+380501112233", addr2, 1, 6)

	//Утворення черги
	manager := NewQueueManager()
	manager.AddClient(client1)
	manager.AddClient(client2)

	//Вивід черги
	fmt.Println("Queue:")
	manager.PrintAll()

	//Сортування клієнтів за місцем у черзі
	manager.Reorder()

	//Вивід черги після оновлення
	fmt.Println("\nAfter updating positions:")
	manager.PrintAll()

	fmt.Printf("Total estimated waiting time: %d minutes\n", manager.TotalEstimatedTime())

	//Серіалізація до nJSON
	fmt.Println("\nJSON Export:")
	fmt.Println(manager.ToJSON())
}
