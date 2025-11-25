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

// Оголошення структури клієнта
type Client struct {
	ID       int
	Name     string
	Email    string
	Phone    string
	Address  Address
	Position int
}

// Функція NewClient створює нового клієнта
// Position відповідає за місце клієнтв в черзі. Задається іншою функцією
func NewClient(id int, name, email, phone string, address Address) Client {
	return Client{
		ID:       id,
		Name:     name,
		Email:    email,
		Phone:    phone,
		Address:  address,
		Position: 0,
	}
}

// Функція Info виводить інформацію клієнта
func (c Client) Info() string {
	return fmt.Sprintf("ID: %d\nName: %s\nEmail: %s\nPhone: %s\nCity: %s\nStreet: %s\nZip: %s\nPosition in queue: %d",
		c.ID, c.Name, c.Email, c.Phone, c.Address.City, c.Address.Street, c.Address.ZipCode, c.Position)
}

// Функція UpdatePosition задає нову позицію клієнта
func (c *Client) UpdatePosition(newPos int) {
	c.Position = newPos
}

func main() {
	fmt.Printf("E-Queue\n\n")

	//Створення адреси
	addr1 := Address{City: "Kyiv", Street: "Khreshchatyk", ZipCode: "01001"}

	//Створення клієнта через структуру addr1
	client1 := NewClient(1, "Ivan Ivanov", "ivan@example.com", "+380501234567", addr1)

	//Створення клієнта через створення нової структури адреси у ньому
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

	fmt.Println("Client 1 info:\n", client1.Info())
	fmt.Println("Client 2 info:\n", client2.Info())

	//Задання місця в черзі клієнтові
	client1.UpdatePosition(5)
	client2.UpdatePosition(3)

	fmt.Println("\nAfter updating positions:")
	fmt.Println("Client 1 info:\n", client1.Info())
	fmt.Println("Client 2 info:\n", client2.Info())
}
