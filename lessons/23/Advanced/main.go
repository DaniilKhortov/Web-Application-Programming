package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Post — структура, яка моделює запит і відповідь API
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

	// Створюємо Go-структуру для надсилання в POST-запиті
	newPost := Post{
		UserID: 101,
		Title:  "New client in queue",
		Body:   "Client was successfully registered into queue",
	}

	// Серіалізуємо структуру у JSON
	jsonData, err := json.Marshal(newPost)
	if err != nil {
		fmt.Println("JSON serialization error:", err)
		return
	}

	// Створюємо власний HTTP-клієнт із таймаутом 5 секунд
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Формуємо POST-запит з JSON-тілом
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Request creation error:", err)
		return
	}

	// Встановлюємо заголовок Content-Type
	req.Header.Set("Content-Type", "application/json")

	// Виконуємо запит
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	// Перевіряємо статус відповіді
	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Unexpected status: %d %s\n", resp.StatusCode, resp.Status)
		return
	}

	// Десеріалізуємо відповідь від сервера у структуру
	var createdPost Post
	err = json.NewDecoder(resp.Body).Decode(&createdPost)
	if err != nil {
		fmt.Println("Error ocured during answer decoding:", err)
		return
	}

	// Виводимо результат
	fmt.Println("Data were successfully sent to server!")
	fmt.Printf("Recieved answer:\n")
	fmt.Printf("ID: %d\n", createdPost.ID)
	fmt.Printf("UserID: %d\n", createdPost.UserID)
	fmt.Printf("Title: %s\n", createdPost.Title)
	fmt.Printf("Content: %s\n", createdPost.Body)
}
