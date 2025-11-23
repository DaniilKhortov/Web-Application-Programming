package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// URL публічного API — у реальному застосунку це був би ендпоїнт черги
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Виконання GET-запиту
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	// Обов’язково закриваємо тіло відповіді
	defer resp.Body.Close()

	// Перевірка статус-коду
	if resp.StatusCode == http.StatusOK {
		// Зчитування тіла відповіді
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Answer parse error:", err)
			return
		}
		// Вивід результату
		fmt.Println("Receiwing answer from server:")
		fmt.Println(string(body))
	} else {
		fmt.Printf("Unexpected code: %d %s\n", resp.StatusCode, resp.Status)
	}
}
