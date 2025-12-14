package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Структура конфігурації застосунку
type Config struct {
	Env          string
	Port         string
	ServiceName  string
	MaxQueueSize int
}

// Єдина функція завантаження конфігурації
func LoadConfig() Config {
	env := getEnv("GO_ENV", "development")

	// Завантаження .env залежно від середовища
	envFile := ".env." + env
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("Failed to load %s, using default parameters\n", envFile)
	}

	return Config{
		Env:          env,
		Port:         getEnv("APP_PORT", "8080"),
		ServiceName:  getEnv("SERVICE_NAME", "Queue Service"),
		MaxQueueSize: getEnvAsInt("MAX_QUEUE_SIZE", 50),
	}
}

// Значення за замовчуванням (string)
func getEnv(key, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}

// Значення за замовчуванням (int)
func getEnvAsInt(key string, defaultVal int) int {
	if value, err := strconv.Atoi(os.Getenv(key)); err == nil {
		return value
	}
	return defaultVal
}
