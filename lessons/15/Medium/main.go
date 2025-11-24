package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func copyLines(srcFile, dstFile string) (int, error) {

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

	scanner := bufio.NewScanner(src)
	writer := bufio.NewWriter(dst)
	defer writer.Flush()

	lineCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return lineCount, fmt.Errorf("error writing to destination file: %v", err)
		}
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return lineCount, fmt.Errorf("error reading source file: %v", err)
	}

	return lineCount, nil
}

func main() {
	filename := "queue.txt"
	copyFile := "queue_copy.txt"

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

	srcFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file for reading:", err)
		return
	}
	defer srcFile.Close()

	scanner := bufio.NewScanner(srcFile)
	lineCount := 0
	fmt.Println("\nFile content (line by line):")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error during scanning file:", err)
		return
	}

	fmt.Printf("\nTotal lines read: %d\n", lineCount)

	copiedLines, err := copyLines(filename, copyFile)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}
	fmt.Printf("Data copied successfully to %s (%d lines)\n", copyFile, copiedLines)
}
