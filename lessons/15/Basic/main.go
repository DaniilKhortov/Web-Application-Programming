package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//Зберігаємо шлях текстового файлу
	filename := "queue.txt"

	//Створюємо файл
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error occured during file creation:", err)
		return
	}

	//Закриваємо файл по закінченю всіх операцій
	defer file.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter client data (Name Surname ID).")

	//Зчитування даних клієнта з консолі
	for {
		fmt.Print("Client: ")
		text, _ := reader.ReadString('\n')

		trimmedText := strings.TrimSpace(text)

		//Умова завершення вводу
		if trimmedText == "" {
			break
		}

		//Запис даних до файлу
		_, err := file.WriteString(text)
		if err != nil {
			fmt.Println("File writing error:", err)
			return
		}
	}

	fmt.Println("\nData was saved successfully:", filename)

	//Зчитування файлу
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error:", err)
		return
	}

	//Вивід вмісту файлу
	fmt.Println("\nFile content:")
	fmt.Println(string(content))
}
