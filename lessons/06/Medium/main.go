package main

import (
	"fmt"
)

// Структура клієнта
type Client struct {
	Name  string
	Email *string
}

// Вузол зв'язаного списку
type Node struct {
	Data *Client
	Next *Node
}

// Зв'язаний список черги
type LinkedQueue struct {
	Head  *Node
	Tail  *Node
	Count int
}

// Функція newClient створює нового клієнта
func newClient(name string, email *string) *Client {
	//Повертаємо адресу памя'тті на утворений об'єкт
	return &Client{Name: name, Email: email}
}

// Функція Enqueue реєструє клієнта у чергу
func (q *LinkedQueue) Enqueue(c *Client) {

	node := &Node{Data: c}

	//Додавання клієнта у кінець зв'язаного списку
	//Виділення місця в списку шляхом перевизначення кінця списку новим значенням
	if q.Tail != nil {
		q.Tail.Next = node
	} else {
		q.Head = node
	}

	q.Tail = node
	q.Count++
}

// Функція Dequeue видаляє клієнта з черги
func (q *LinkedQueue) Dequeue() *Client {

	//Перевірка на наявність значень у списку
	if q.Head == nil {
		return nil
	}

	//Прибирання клієнта зі списку шляхом перевизначення початку списку
	client := q.Head.Data
	q.Head = q.Head.Next
	if q.Head == nil {
		q.Tail = nil
	}
	q.Count--
	return client
}

// Функція Find шукає клієнта за ім'ям
func (q *LinkedQueue) Find(name string) *Client {
	for current := q.Head; current != nil; current = current.Next {
		if current.Data.Name == name {
			return current.Data
		}
	}
	return nil
}

func main() {
	fmt.Println("E-Queue - Register desk")

	fmt.Println("\nPointer demonstration")
	fmt.Println("-----------------------")

	//totalClients - зберігає кількість клієнтів
	var totalClients int = 0

	//pTotal - вказівник на totalClients
	pTotal := &totalClients

	fmt.Println("totalClients address:", pTotal)
	fmt.Println("totalClients value:", *pTotal)

	//СТворення клієнтів
	email1 := "olena@example.com"
	client1 := newClient("Olena", &email1)
	client2 := newClient("Mykhailo", nil)
	client3 := newClient("Fedir", nil)

	//Ініціалізація черги як зв`язний список
	var queue LinkedQueue

	fmt.Println("\nRegistering clients")
	fmt.Println("-----------------------")

	//Додавання клієнтів
	queue.Enqueue(client1)
	queue.Enqueue(client2)
	queue.Enqueue(client3)
	fmt.Printf("Total clients in queue: %d\n", queue.Count)

	//Обслуговування клієнтів
	fmt.Println("\nServing clients:")
	fmt.Println("-----------------------")
	for i := 0; i < 2; i++ {
		c := queue.Dequeue()
		if c != nil {
			if c.Email != nil {
				fmt.Printf("Serving %s (%s)\n", c.Name, *c.Email)
			} else {
				fmt.Printf("Serving %s (no email)\n", c.Name)
			}
		}
	}

	fmt.Printf("Clients left: %d\n", queue.Count)

	//Пошук клієнта, який існує у черзі
	fmt.Println("\nSearching for 'Fedir':")
	found := queue.Find("Fedir")
	if found != nil {
		fmt.Printf("Found client: %s\n", found.Name)
	} else {
		fmt.Println("Client not found")
	}

	fmt.Println("\nUsing new()")
	fmt.Println("-----------------------")

	//Обслуговування середнього клієнта в списку
	pAverage := new(float64)
	if pAverage != nil {
		*pAverage = float64(*pTotal) / 2.0
		fmt.Println("Average client served:", *pAverage)
	}

	//Перевірка на наявність значення pNil
	fmt.Println("\nTesting nil-pointer")
	fmt.Println("-----------------------")
	var pNil *int
	if pNil == nil {
		fmt.Println("Nil pointer exception!")
	}
}
