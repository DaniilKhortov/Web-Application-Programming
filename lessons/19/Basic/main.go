package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// --- 1. Підключення до бази даних ---
	// Якщо ти в XAMPP і пароль порожній — прибери "1234"
	connStr := "root:@tcp(127.0.0.1:3306)/queue_db"

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("Connection error: %v\n", err)
	}
	defer db.Close()

	// Перевірка з'єднання
	err = db.Ping()
	if err != nil {
		log.Fatalf("No connection with database: %v\n", err)
	}
	fmt.Println("Connection established!")

	// --- 2. Створення таблиці ---
	createTable := `
	CREATE TABLE IF NOT EXISTS queue (
		id INT AUTO_INCREMENT PRIMARY KEY,
		client_name VARCHAR(100),
		service_type VARCHAR(50),
		ticket_number INT
	);
	`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatalf("Error ocured during table creation: %v\n", err)
	}
	fmt.Println("Table 'queue' successfully created!")

	// --- 3. INSERT-запит ---
	insertQuery := `INSERT INTO queue (client_name, service_type, ticket_number) VALUES (?, ?, ?)`

	res, err := db.Exec(insertQuery, "Ivan Petrenko", "Communal services payment", 101)
	if err != nil {
		log.Fatalf("Insertion error: %v\n", err)
	}

	newID, _ := res.LastInsertId()
	fmt.Printf("Added new item with ID = %d\n", newID)

	// --- 4. SELECT-запит ---
	var (
		clientName   string
		serviceType  string
		ticketNumber int
	)

	selectQuery := `SELECT client_name, service_type, ticket_number FROM queue WHERE id = ?`
	err = db.QueryRow(selectQuery, newID).Scan(&clientName, &serviceType, &ticketNumber)
	if err != nil {
		log.Fatalf("Reading error: %v\n", err)
	}

	fmt.Println("Received data:")
	fmt.Printf("Client: %s | Service: %s | Ticket: %d\n",
		clientName, serviceType, ticketNumber)
}
