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

// Основна структура: Клієнт електронної черги
type Client struct {
	ID       int
	Name     string
	Email    string
	Phone    string
	Address  Address
	Position int // Позиція в черзі
}

// Конструктор для Client
func NewClient(id int, name, email, phone string, address Address) Client {
	return Client{
		ID:       id,
		Name:     name,
		Email:    email,
		Phone:    phone,
		Address:  address,
		Position: 0, // Спочатку позиція = 0
	}
}

// Метод з value receiver для читання даних
func (c Client) Info() string {
	return fmt.Sprintf("ID: %d\nName: %s\nEmail: %s\nPhone: %s\nCity: %s\nStreet: %s\nZip: %s\nPosition in queue: %d",
		c.ID, c.Name, c.Email, c.Phone, c.Address.City, c.Address.Street, c.Address.ZipCode, c.Position)
}

// Метод з pointer receiver для модифікації даних
func (c *Client) UpdatePosition(newPos int) {
	c.Position = newPos
}

func main() {
	fmt.Printf("E-Queue\n\n")
	// Ініціалізація структури через конструктор
	addr1 := Address{City: "Kyiv", Street: "Khreshchatyk", ZipCode: "01001"}
	client1 := NewClient(1, "Ivan Ivanov", "ivan@example.com", "+380501234567", addr1)

	// Ініціалізація без конструктора
	client2 := Client{
		ID:    2,
		Name:  "Olena Petrova",
		Email: "olena@example.com",
		Phone: "+380501112233",
		Address: Address{
			City:    "Lviv",
			Street:  "Svobody",
			ZipCode: "79000",
		},
		Position: 0,
	}

	// Виклик value receiver методу
	fmt.Println("Client 1 info:\n", client1.Info())
	fmt.Println("Client 2 info:\n", client2.Info())

	// Модифікація через pointer receiver
	client1.UpdatePosition(5)
	client2.UpdatePosition(3)

	fmt.Println("\nAfter updating positions:")
	fmt.Println("Client 1 info:\n", client1.Info())
	fmt.Println("Client 2 info:\n", client2.Info())
}
