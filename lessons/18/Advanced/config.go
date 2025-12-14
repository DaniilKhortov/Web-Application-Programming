package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Структура конфігурації
type Config struct {
	Env          string
	Port         string
	ServiceName  string
	MaxQueueSize int
}

// Головна функція завантаження конфігурації
func LoadConfig() (Config, error) {
	env := getEnv("GO_ENV", "development")

	// Завантаження файлу відповідного середовища
	loadEnvFiles(env)

	cfg := Config{
		Env:          env,
		Port:         getEnv("APP_PORT", "8080"),
		ServiceName:  getEnv("SERVICE_NAME", "Queue Service"),
		MaxQueueSize: getEnvAsInt("MAX_QUEUE_SIZE", 50),
	}

	// Валідація конфігурації
	if err := validateConfig(cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

// Пріоритет завантаження файлів
func loadEnvFiles(env string) {
	files := []string{
		// базовий (опційний)
		".env",
		// середовище
		".env." + env,
	}

	for _, file := range files {
		if err := godotenv.Overload(file); err == nil {
			log.Printf("Loaded: %s\n", file)
		}
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

// Функція валідації
func validateConfig(cfg Config) error {
	if cfg.Port == "" {
		return errors.New("APP_PORT cannot be empty")
	}

	port, err := strconv.Atoi(cfg.Port)
	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("invalid APP_PORT: %s", cfg.Port)
	}

	if cfg.MaxQueueSize <= 0 {
		return errors.New("MAX_QUEUE_SIZE must be > 0")
	}

	if cfg.ServiceName == "" {
		return errors.New("SERVICE_NAME not defined")
	}

	return nil
}
