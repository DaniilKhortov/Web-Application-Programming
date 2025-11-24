package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	filename := "queue.txt"

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error occured during file creation:", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter client data (Name Surname ID).")

	for {
		fmt.Print("Client: ")
		text, _ := reader.ReadString('\n')

		trimmedText := strings.TrimSpace(text)

		if trimmedText == "" {
			break
		}

		_, err := file.WriteString(text)
		if err != nil {
			fmt.Println("File writing error:", err)
			return
		}
	}

	fmt.Println("\nData was saved successfully:", filename)

	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error:", err)
		return
	}

	fmt.Println("\nFile content:")
	fmt.Println(string(content))
}
