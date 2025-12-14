package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Структура для представлення клієнта в черзі
type QueueClient struct {
	ID           int
	ClientName   string
	ServiceType  string
	TicketNumber int
}

// Функція для додавання нового клієнта в чергу
func addClient(db *sql.DB, clientName, serviceType string, ticketNumber int) error {
	insertQuery := `INSERT INTO queue (client_name, service_type, ticket_number) VALUES (?, ?, ?)`

	result, err := db.Exec(insertQuery, clientName, serviceType, ticketNumber)
	if err != nil {
		return fmt.Errorf("insertion error: %v", err)
	}

	// Отримання ID доданого запису
	lastID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error occurred during recieving id: %v", err)
	}

	fmt.Printf("Added new item with ID = %d, ticket №%d\n", lastID, ticketNumber)
	return nil
}

// Функція для отримання всіх клієнтів з черги
func getAllClients(db *sql.DB) ([]QueueClient, error) {
	selectQuery := `SELECT id, client_name, service_type, ticket_number FROM queue`

	// Виконання запиту для отримання кількох рядків
	rows, err := db.Query(selectQuery)
	if err != nil {
		return nil, fmt.Errorf("error occurred during request: %v", err)
	}
	defer rows.Close()

	// Створення зрізу для зберігання клієнтів
	var clients []QueueClient

	// Ітерація по всіх рядках результату
	for rows.Next() {
		var client QueueClient

		// Сканування даних з рядка в структуру
		err := rows.Scan(&client.ID, &client.ClientName, &client.ServiceType, &client.TicketNumber)
		if err != nil {
			return nil, fmt.Errorf("raw scanning error: %v", err)
		}

		// Додавання клієнта до зрізу
		clients = append(clients, client)
	}

	// Перевірка на помилки, що могли виникнути під час ітерації
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error occured during iteration: %v", err)
	}

	return clients, nil
}

// Функція для ініціалізації бази даних (створення таблиці)
func initDatabase(db *sql.DB) error {
	createTable := `
	CREATE TABLE IF NOT EXISTS queue (
		id INT AUTO_INCREMENT PRIMARY KEY,
		client_name VARCHAR(100) NOT NULL,
		service_type VARCHAR(50) NOT NULL,
		ticket_number INT NOT NULL UNIQUE
	);
	`

	_, err := db.Exec(createTable)
	if err != nil {
		return fmt.Errorf("error occurred during table creation: %v", err)
	}

	fmt.Println("Table 'queue' successfully created!")
	return nil
}

// Функція для виведення списку клієнтів
func displayClients(clients []QueueClient) {

	fmt.Println("E-Queue")

	if len(clients) == 0 {
		fmt.Println("Queue is empty")
	} else {
		for _, client := range clients {
			fmt.Printf("%-5d | %-25s | %-20s | %-10d\n",
				client.ID,
				client.ClientName,
				client.ServiceType,
				client.TicketNumber)
		}
	}

	fmt.Printf("Clients in queue: %d\n\n", len(clients))
}

func main() {
	// Підключення до бази даних MySQL
	connStr := "root:@tcp(127.0.0.1:3306)/queue_db"

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("Connection error: %v\n", err)
	}
	defer db.Close()

	// Перевірка з'єднання з базою даних
	err = db.Ping()
	if err != nil {
		log.Fatalf("No connection with database: %v\n", err)
	}
	fmt.Println("Connection with database established!")

	// Ініціалізація бази даних (створення таблиці)
	err = initDatabase(db)
	if err != nil {
		log.Fatalf("Error %v\n", err)
	}

	fmt.Println("\nAdding clients to queue")

	// Додавання кількох клієнтів до черги
	clientsToAdd := []struct {
		name    string
		service string
		ticket  int
	}{
		{"Ivan Petrenko", "Communal services payment", 101},
		{"Maria Kovalenko", "Consultation", 102},
		{"Oleksandr Shevchenko", "Reportong counters data", 103},
		{"Natalia Bondarenko", "Communal services payment", 104},
		{"Dmytro Melnyk", "Reportong counters data", 105},
	}

	for _, client := range clientsToAdd {
		err := addClient(db, client.name, client.service, client.ticket)
		if err != nil {
			log.Printf("Operation error: %v\n", err)
		}
	}

	fmt.Println("\nRecieving all client data from database")

	// Отримання всіх клієнтів з бази даних
	clients, err := getAllClients(db)
	if err != nil {
		log.Fatalf("Connection with clients data error: %v\n", err)
	}

	// Виведення списку клієнтів
	displayClients(clients)
}
