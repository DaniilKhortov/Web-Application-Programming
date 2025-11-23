package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Функція для копіювання файлу блоками та підрахунку байтів
func copyLargeFile(srcFile, dstFile string, bufSize int) (int64, error) {
	src, err := os.Open(srcFile)
	if err != nil {
		return 0, fmt.Errorf("cannot open source file: %v", err)
	}
	defer src.Close()

	dst, err := os.Create(dstFile)
	if err != nil {
		return 0, fmt.Errorf("cannot create destination file: %v", err)
	}
	defer dst.Close()

	reader := bufio.NewReader(src)
	writer := bufio.NewWriter(dst)
	defer writer.Flush()

	var totalBytes int64 = 0
	buffer := make([]byte, bufSize)

	start := time.Now()
	for {
		n, err := reader.Read(buffer)
		if n > 0 {
			totalBytes += int64(n)
			if _, wErr := writer.Write(buffer[:n]); wErr != nil {
				return totalBytes, fmt.Errorf("write error: %v", wErr)
			}
		}
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return totalBytes, fmt.Errorf("read error: %v", err)
		}
	}
	duration := time.Since(start)
	fmt.Printf("Copied %d bytes in %v\n", totalBytes, duration)
	return totalBytes, nil
}

func main() {
	filename := "queue.txt"
	copyFile := "queue_large_copy.txt"
	bufSize := 1024 * 16 // 16 KB блок

	// Введення даних користувачем
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter client data (Name Surname ID). Press Enter on empty line to finish:")

	for {
		fmt.Print("Client: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			break
		}

		_, err := file.WriteString(text + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("\nData saved successfully to", filename)

	// Читання великих файлів блочно
	fmt.Println("\nStarting to copy large file...")
	totalBytes, err := copyLargeFile(filename, copyFile, bufSize)
	if err != nil {
		fmt.Println("Error copying large file:", err)
		return
	}
	fmt.Printf("Large file copy completed. Total bytes: %d\n", totalBytes)
}
