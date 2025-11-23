package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Post — структура, яка відповідає JSON-відповіді з API
// https://jsonplaceholder.typicode.com/posts/1
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// getPost виконує GET-запит до API, зчитує JSON і перетворює його у структуру Post
func getPost(url string) (*Post, error) {
	// Виконання GET-запиту
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	// Перевірка статус-коду відповіді
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected code: %d %s", resp.StatusCode, resp.Status)
	}

	// Десеріалізація JSON у структуру
	var post Post
	err = json.NewDecoder(resp.Body).Decode(&post)
	if err != nil {
		return nil, fmt.Errorf("error ocured during JSON decode: %v", err)
	}

	return &post, nil
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Викликаємо функцію отримання даних
	post, err := getPost(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Виводимо окремі поля структури
	fmt.Println("Data successfully receiwed")
	fmt.Printf("ID of publication: %d\n", post.ID)
	fmt.Printf("Author (UserID): %d\n", post.UserID)
	fmt.Printf("Title: %s\n", post.Title)
	fmt.Printf("Content:\n%s\n", post.Body)
}
