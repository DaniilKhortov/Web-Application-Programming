package main

import (
	"fmt"
)

// Вкладена структура: Адреса клієнта
type Address struct {
	City    string
	Street  string
	ZipCode string
}

// Функція-конструктор для Address
func NewAddress(city, street, zip string) Address {
	return Address{
		City:    city,
		Street:  street,
		ZipCode: zip,
	}
}

// Додаткова структура: Дані черги
type QueueData struct {
	Position int
	AvgTime  int // Середній час обслуговування одного клієнта (хвилини)
}

// Метод обчислення похідного значення (для черги)
func (q QueueData) EstimatedWait() int {
	return q.Position * q.AvgTime
}

// Основна структура: Клієнт електронної черги
// Використовує embedding для QueueData (анонімне поле)
type Client struct {
	ID        int
	Name      string
	Email     string
	Phone     string
	Address   Address
	QueueData // анонімне поле
}

// Конструктор для Client
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

// Метод з value receiver — читання даних
func (c Client) Info() string {
	return fmt.Sprintf(
		"Client #%d: %s\nEmail: %s | Phone: %s\nAddress: %s, %s (%s)\nPosition: %d | Wait time: %d min\n",
		c.ID, c.Name, c.Email, c.Phone, c.Address.City, c.Address.Street, c.Address.ZipCode,
		c.Position, c.EstimatedWait(),
	)
}

// Метод з pointer receiver — зміна позиції в черзі
func (c *Client) UpdatePosition(newPos int) {
	c.Position = newPos
}

// Метод з pointer receiver — зміна середнього часу
func (c *Client) UpdateAvgTime(newAvg int) {
	c.AvgTime = newAvg
}

func main() {
	fmt.Printf("E-Queue\n\n")

	// Створення об'єктів через функції-конструктори
	addr1 := NewAddress("Kyiv", "Khreshchatyk", "01001")
	client1 := NewClient(1, "Ivan Ivanov", "ivan@example.com", "+380501234567", addr1, 4, 3)

	addr2 := NewAddress("Lviv", "Svobody Ave", "79000")
	client2 := NewClient(2, "Olena Petrova", "olena@example.com", "+380501112233", addr2, 2, 5)

	// Вивід інформації (value receiver)
	fmt.Println("Client 1 info:\n", client1.Info())
	fmt.Println("Client 2 info:\n", client2.Info())

	// Оновлення через pointer receiver
	client1.UpdatePosition(1)
	client2.UpdateAvgTime(6)

	fmt.Println("\nAfter updating positions:")
	fmt.Println("Client 1 info:\n", client1.Info())
	fmt.Println("Client 2 info:\n", client2.Info())
}
