package main

import (
	"fmt"
)

// Оголошення структури адреси
type Address struct {
	City    string
	Street  string
	ZipCode string
}

// Функція NewAddress створює та повертає адресу
func NewAddress(city, street, zip string) Address {
	return Address{
		City:    city,
		Street:  street,
		ZipCode: zip,
	}
}

// Оголошення структури черги
type QueueData struct {
	Position int
	AvgTime  int
}

// Функція EstimatedWait розраховує приблизний час очікування
func (q QueueData) EstimatedWait() int {
	return q.Position * q.AvgTime
}

// Оголошення структури клієнта
type Client struct {
	ID      int
	Name    string
	Email   string
	Phone   string
	Address Address
	//QueueData - це анонімне поле. Вона вбудовує поля та методи QueueData до Client
	//Таким чином утворюється композиція
	QueueData
}

// Функція NewClient створює нового клієнта
func NewClient(id int, name, email, phone string, addr Address, pos int, avg int) Client {
	return Client{
		ID:        id,
		Name:      name,
		Email:     email,
		Phone:     phone,
		Address:   addr,
		QueueData: QueueData{Position: pos, AvgTime: avg},
	}
}

// Функція Info виводить інформацію клієнта
func (c Client) Info() string {
	return fmt.Sprintf(
		"Client #%d: %s\nEmail: %s | Phone: %s\nAddress: %s, %s (%s)\nPosition: %d | Wait time: %d min\n",
		c.ID, c.Name, c.Email, c.Phone, c.Address.City, c.Address.Street, c.Address.ZipCode,
		c.Position, c.EstimatedWait(),
	)
}

// Функція UpdatePosition задає нову позицію клієнта
func (c *Client) UpdatePosition(newPos int) {
	c.Position = newPos
}

// Функція UpdateAvgTime задає новий час очікування для клієнта
func (c *Client) UpdateAvgTime(newAvg int) {
	c.AvgTime = newAvg
}

func main() {
	fmt.Printf("E-Queue\n\n")

	//Створення адрес та клієнтів
	addr1 := NewAddress("Kyiv", "Khreshchatyk", "01001")
	client1 := NewClient(1, "Ivan Ivanov", "ivan@example.com", "+380501234567", addr1, 4, 3)

	addr2 := NewAddress("Lviv", "Svobody Ave", "79000")
	client2 := NewClient(2, "Olena Petrova", "olena@example.com", "+380501112233", addr2, 2, 5)

	fmt.Println("Client 1 info:\n", client1.Info())
	fmt.Println("Client 2 info:\n", client2.Info())

	//Першому клієнту задаємо позицію в черзі
	client1.UpdatePosition(1)

	//Другому клієнту задаємо час очікування
	client2.UpdateAvgTime(6)

	fmt.Println("\nAfter updating positions:")
	fmt.Println("Client 1 info:\n", client1.Info())
	fmt.Println("Client 2 info:\n", client2.Info())
}
