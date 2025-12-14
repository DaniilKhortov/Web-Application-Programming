package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//Підключення до бази даних (MySQL)
	connStr := "root:@tcp(127.0.0.1:3306)/queue_db"

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("Connection error: %v\n", err)
	}

	//Закриття доступу до бази даних після завершення всіх операцій
	defer db.Close()

	//Перевірка на підключення до бази даних
	err = db.Ping()
	if err != nil {
		log.Fatalf("No connection with database: %v\n", err)
	}
	fmt.Println("Connection established!")

	//Створення SQL (MySQL) запиту для створення таблиці
	createTable := `
	CREATE TABLE IF NOT EXISTS queue (
		id INT AUTO_INCREMENT PRIMARY KEY,
		client_name VARCHAR(100),
		service_type VARCHAR(50),
		ticket_number INT
	);
	`
	//Виконання SQL запиту
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatalf("Error occurred during table creation: %v\n", err)
	}
	fmt.Println("Table 'queue' successfully created!")

	//Створення запиту на додання запису до таблиці
	//Символ ? - це розмітка для аргументів
	insertQuery := `INSERT INTO queue (client_name, service_type, ticket_number) VALUES (?, ?, ?)`

	//Виконання SQL запиту
	//Параметри необхідно надати при надсиланні запиту. Порядок аргументів визначає місце в запиті
	res, err := db.Exec(insertQuery, "Ivan Petrenko", "Communal services payment", 101)
	if err != nil {
		log.Fatalf("Insertion error: %v\n", err)
	}

	//Отримання ідентифікатора останнього запису для перевірки бази даних
	newID, _ := res.LastInsertId()
	fmt.Printf("Added new item with ID = %d\n", newID)

	//Створення структури клієнтів
	var (
		clientName   string
		serviceType  string
		ticketNumber int
	)

	//Створення запиту отримання запису з бази даних
	selectQuery := `SELECT client_name, service_type, ticket_number FROM queue WHERE id = ?`

	//Виконання SQL запиту отримання запису
	err = db.QueryRow(selectQuery, newID).Scan(&clientName, &serviceType, &ticketNumber)
	if err != nil {
		log.Fatalf("Reading error: %v\n", err)
	}

	//Вивід результатів
	fmt.Println("Received data:")
	fmt.Printf("Client: %s | Service: %s | Ticket: %d\n",
		clientName, serviceType, ticketNumber)
}
