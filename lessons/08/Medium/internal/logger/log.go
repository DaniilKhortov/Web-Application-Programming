package logger

import "fmt"

// logInternal — приватна функція, недоступна ззовні.
func logInternal(msg string) {
	fmt.Printf("[internal] %s\n", msg)
}

// LogInfo — експортована функція для загальних логів.
func LogInfo(msg string) {
	logInternal(msg)
}
