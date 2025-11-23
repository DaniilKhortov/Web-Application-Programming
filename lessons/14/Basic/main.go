package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Назва файлу для збереження черги
	filename := "queue.txt"

	// Створення або відкриття файлу для запису
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error occured during file creation:", err)
		return
	}
	// Гарантоване закриття файлу після завершення
	defer file.Close()

	// Створення буферизованого введення з консолі
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter client data (Name Surname ID).")

	for {
		fmt.Print("Client: ")
		text, _ := reader.ReadString('\n')

		// 2. Очищаємо рядок від пробілів на початку/в кінці та символу \n
		trimmedText := strings.TrimSpace(text)

		// 3. Перевіряємо, чи рядок порожній
		if trimmedText == "" {
			break // Якщо так — виходимо з циклу
		}

		// Запис у файл потрібно робити з оригінальним рядком, щоб зберегти перехід на новий рядок
		_, err := file.WriteString(text)
		if err != nil {
			fmt.Println("File writing error:", err)
			return
		}
	}

	fmt.Println("\nData was saved successfully:", filename)

	// Зчитування всього вмісту файлу
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error:", err)
		return
	}

	// Виведення даних на екран
	fmt.Println("\nFile content:")
	fmt.Println(string(content))
}
