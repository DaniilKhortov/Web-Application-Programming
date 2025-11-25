package logger

import "fmt"

//Функція logInternal обгортає повідомлення для кращого виводу
func logInternal(msg string) {
	fmt.Printf("[internal] %s\n", msg)
}

//Функція LogInfo логує дію
func LogInfo(msg string) {
	logInternal(msg)
}
