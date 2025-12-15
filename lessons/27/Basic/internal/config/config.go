package config

import (
	"os"
)

type Config struct {
	AppName string
	Port    string
	Env     string
}

func getEnv(key, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}

// Функція завантаження файлу конфігурації
func Load() Config {
	return Config{
		AppName: getEnv("APP_NAME", "queue-app"),
		Port:    getEnv("APP_PORT", "8080"),
		Env:     getEnv("APP_ENV", "development"),
	}
}
