package main

import (
	"os"
	"strconv"
)

// Структура конфігурації
type Config struct {
	Port         string
	ServiceName  string
	MaxQueueSize int
}

// Функція з значеннями за замовчуванням
func LoadConfig() Config {
	return Config{
		Port:         getEnv("APP_PORT", "8080"),
		ServiceName:  getEnv("SERVICE_NAME", "Queue Service"),
		MaxQueueSize: getEnvAsInt("MAX_QUEUE_SIZE", 100),
	}
}

// Отримання  середовища типу string
func getEnv(key, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}

// Отримання  середовища типу int
func getEnvAsInt(key string, defaultVal int) int {
	valueStr := os.Getenv(key)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}
