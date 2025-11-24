package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Sensor struct {
	ID       int
	Name     string
	Location string
}

func connectDB() (*sql.DB, error) {
	connStr := "root:@tcp(127.0.0.1:3306)/queue_db"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, fmt.Errorf("connection error: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("no connection with database: %v", err)
	}

	return db, nil
}

func ensureTable(db *sql.DB) error {
	createTable := `
	CREATE TABLE IF NOT EXISTS sensors (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100),
		location VARCHAR(100)
	);`
	_, err := db.Exec(createTable)
	return err
}

func addSensor(db *sql.DB, name string, location string) error {
	query := `INSERT INTO sensors (name, location) VALUES (?, ?)`
	_, err := db.Exec(query, name, location)
	if err != nil {
		return fmt.Errorf("error ocured during insertion: %v", err)
	}
	fmt.Printf("Sensor '%s' (%s) added successfully.\n", name, location)
	return nil
}

func getSensors(db *sql.DB) ([]Sensor, error) {
	query := `SELECT id, name, location FROM sensors`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("SQL error: %v", err)
	}
	defer rows.Close()

	var sensors []Sensor
	for rows.Next() {
		var s Sensor
		if err := rows.Scan(&s.ID, &s.Name, &s.Location); err != nil {
			return nil, fmt.Errorf("reading error: %v", err)
		}
		sensors = append(sensors, s)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error after reading: %v", err)
	}

	return sensors, nil
}

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatalf(" %v", err)
	}
	defer db.Close()
	fmt.Println("Connection established!")

	if err := ensureTable(db); err != nil {
		log.Fatalf("Error ocured during table creation: %v", err)
	}

	_ = addSensor(db, "Heat sensor", "Room №1")
	_ = addSensor(db, "Humidity sensor", "Server room")
	_ = addSensor(db, "Motion sensor", "Entrance")

	sensors, err := getSensors(db)
	if err != nil {
		log.Fatalf(" %v", err)
	}

	fmt.Println("\nSensor list:")
	for _, s := range sensors {
		fmt.Printf("ID: %d | Name: %s | Location: %s\n", s.ID, s.Name, s.Location)
	}
}
