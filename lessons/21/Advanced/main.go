package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Sensor struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

var db *sql.DB

func addSensor(name string, location string) error {
	query := `INSERT INTO sensors (name, location) VALUES (?, ?)`
	_, err := db.Exec(query, name, location)
	return err
}

func getSensors() ([]Sensor, error) {
	rows, err := db.Query(`SELECT id, name, location FROM sensors`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sensors []Sensor
	for rows.Next() {
		var s Sensor
		if err := rows.Scan(&s.ID, &s.Name, &s.Location); err != nil {
			return nil, err
		}
		sensors = append(sensors, s)
	}
	return sensors, nil
}

func handleAddSensor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	var s Sensor
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid format JSON", http.StatusBadRequest)
		return
	}

	if s.Name == "" || s.Location == "" {
		http.Error(w, "Field 'name' and 'location' are obligatory", http.StatusBadRequest)
		return
	}

	if err := addSensor(s.Name, s.Location); err != nil {
		http.Error(w, "Insertion error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"message": "Sensor added"}`)
}

func handleGetSensors(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	sensors, err := getSensors()
	if err != nil {
		http.Error(w, "Receiving error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sensors)
}

func main() {

	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/queue_db"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("SQL error: %v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("No connection: %v", err)
	}
	fmt.Println("Connection established!")

	createTable := `
	CREATE TABLE IF NOT EXISTS sensors (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		location VARCHAR(100) NOT NULL
	);`
	if _, err := db.Exec(createTable); err != nil {
		log.Fatalf("Error ocured during table creation: %v", err)
	}
	fmt.Println("Table 'sensors' successfully created.")

	http.HandleFunc("/sensors", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGetSensors(w, r)
		case http.MethodPost:
			handleAddSensor(w, r)
		default:
			http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server launched at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
