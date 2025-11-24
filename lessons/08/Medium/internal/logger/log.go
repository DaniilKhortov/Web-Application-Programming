package logger

import "fmt"

func logInternal(msg string) {
	fmt.Printf("[internal] %s\n", msg)
}

func LogInfo(msg string) {
	logInternal(msg)
}
