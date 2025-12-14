package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Структура для представлення клієнта в черзі
type QueueClient struct {
	ID           int    `json:"id"`
	ClientName   string `json:"client_name"`
	ServiceType  string `json:"service_type"`
	TicketNumber int    `json:"ticket_number"`
}

// Структура для валідації вхідних даних при POST-запиті
type CreateClientRequest struct {
	ClientName   string `json:"client_name"`
	ServiceType  string `json:"service_type"`
	TicketNumber int    `json:"ticket_number"`
}

// Структура для стандартизованої відповіді з помилкою
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// Структура для успішної відповіді
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Глобальна змінна для зберігання підключення до БД
var db *sql.DB

// Функція для додавання нового клієнта в чергу
func addClient(clientName, serviceType string, ticketNumber int) (int64, error) {
	insertQuery := `INSERT INTO queue (client_name, service_type, ticket_number) VALUES (?, ?, ?)`

	result, err := db.Exec(insertQuery, clientName, serviceType, ticketNumber)
	if err != nil {
		return 0, fmt.Errorf("insertion error: %v", err)
	}

	// Отримання ID доданого запису
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting last insert ID: %v", err)
	}

	return lastID, nil
}

// Функція для отримання всіх клієнтів з черги
func getAllClients() ([]QueueClient, error) {
	selectQuery := `SELECT id, client_name, service_type, ticket_number FROM queue ORDER BY ticket_number`

	rows, err := db.Query(selectQuery)
	if err != nil {
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()

	var clients []QueueClient

	for rows.Next() {
		var client QueueClient

		err := rows.Scan(&client.ID, &client.ClientName, &client.ServiceType, &client.TicketNumber)
		if err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}

		clients = append(clients, client)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iteration error: %v", err)
	}

	return clients, nil
}

// Функція для ініціалізації бази даних
func initDatabase() error {
	createTable := `
	CREATE TABLE IF NOT EXISTS queue (
		id INT AUTO_INCREMENT PRIMARY KEY,
		client_name VARCHAR(100) NOT NULL,
		service_type VARCHAR(50) NOT NULL,
		ticket_number INT NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(createTable)
	if err != nil {
		return fmt.Errorf("table creation error: %v", err)
	}

	return nil
}

// Функція валідації даних клієнта
func validateClientData(req CreateClientRequest) error {
	// Перевірка обов'язкових полів
	if strings.TrimSpace(req.ClientName) == "" {
		return fmt.Errorf("client_name is required and cannot be empty")
	}

	if strings.TrimSpace(req.ServiceType) == "" {
		return fmt.Errorf("service_type is required and cannot be empty")
	}

	// Перевірка номера квитка
	if req.TicketNumber <= 0 {
		return fmt.Errorf("ticket_number must be a positive integer")
	}

	// Перевірка довжини полів
	if len(req.ClientName) > 100 {
		return fmt.Errorf("client_name must be less than 100 characters")
	}

	if len(req.ServiceType) > 50 {
		return fmt.Errorf("service_type must be less than 50 characters")
	}

	return nil
}

// HTTP обробник для POST-запиту (додавання нового клієнта)
func createClientHandler(w http.ResponseWriter, r *http.Request) {
	// Встановлення заголовка Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Перевірка методу запиту
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "Method Not Allowed",
			Message: "Only POST method is allowed",
		})
		return
	}

	// Декодування JSON з тіла запиту
	var req CreateClientRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "Bad Request",
			Message: fmt.Sprintf("Invalid JSON format: %v", err),
		})
		return
	}

	// Валідація даних
	if err := validateClientData(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "Validation Error",
			Message: err.Error(),
		})
		return
	}

	// Додавання клієнта до бази даних
	clientID, err := addClient(req.ClientName, req.ServiceType, req.TicketNumber)
	if err != nil {
		// Перевірка на дублікат номера квитка
		if strings.Contains(err.Error(), "Duplicate entry") {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{
				Error:   "Duplicate Ticket Number",
				Message: fmt.Sprintf("Ticket number %d already exists", req.TicketNumber),
			})
			return
		}

		// Інші помилки сервера
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "Internal Server Error",
			Message: fmt.Sprintf("Failed to add client: %v", err),
		})
		return
	}

	// Успішна відповідь
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "Client added successfully",
		Data: map[string]interface{}{
			"id":            clientID,
			"client_name":   req.ClientName,
			"service_type":  req.ServiceType,
			"ticket_number": req.TicketNumber,
		},
	})

	log.Printf("New client added: %s (Ticket #%d, ID: %d)", req.ClientName, req.TicketNumber, clientID)
}

// HTTP обробник для GET-запиту (отримання всіх клієнтів)
func getClientsHandler(w http.ResponseWriter, r *http.Request) {
	// Встановлення заголовка Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Перевірка методу запиту
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "Method Not Allowed",
			Message: "Only GET method is allowed",
		})
		return
	}

	// Отримання списку клієнтів з бази даних
	clients, err := getAllClients()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "Internal Server Error",
			Message: fmt.Sprintf("Failed to retrieve clients: %v", err),
		})
		return
	}

	// Успішна відповідь
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "Clients retrieved successfully",
		Data: map[string]interface{}{
			"clients": clients,
			"count":   len(clients),
		},
	})

	log.Printf("✓ Retrieved %d clients from queue", len(clients))
}

// Обробник для кореневого шляху (інформація про API)
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	apiInfo := map[string]interface{}{
		"service": "E-Queue Management System",
		"version": "1.0.0",
		"endpoints": map[string]string{
			"POST /clients": "Add a new client to the queue",
			"GET /clients":  "Get all clients in the queue",
		},
		"example_post": map[string]interface{}{
			"client_name":   "Ivan Petrenko",
			"service_type":  "Communal services payment",
			"ticket_number": 101,
		},
	}

	json.NewEncoder(w).Encode(apiInfo)
}

func main() {
	// Підключення до бази даних MySQL
	connStr := "root:@tcp(127.0.0.1:3306)/queue_db"

	var err error
	db, err = sql.Open("mysql", connStr)
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

	// Ініціалізація бази даних
	err = initDatabase()
	if err != nil {
		log.Fatalf("Database initialization error: %v\n", err)
	}
	fmt.Println("Database initialized successfully!")

	// Налаштування HTTP маршрутів
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/clients", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createClientHandler(w, r)
		case http.MethodGet:
			getClientsHandler(w, r)
		default:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(ErrorResponse{
				Error:   "Method Not Allowed",
				Message: "Only GET and POST methods are allowed",
			})
		}
	})

	// Запуск HTTP-сервера
	port := ":8080"
	fmt.Printf("\nREST API Server is running on http://localhost%s\n", port)
	fmt.Println("\nAvailable endpoints:")
	fmt.Println("  GET  http://localhost:8080/        - API information")
	fmt.Println("  POST http://localhost:8080/clients - Add new client")
	fmt.Println("  GET  http://localhost:8080/clients - Get all clients")

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server error: %v\n", err)
	}
}
