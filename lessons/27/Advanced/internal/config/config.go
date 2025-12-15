package config

import (
	"errors"
	"os"
)

type Config struct {
	AppName string
	Port    string
	Env     string
}

func getEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}

// Функція завантаження файлу конфігурації
func Load() (Config, error) {
	cfg := Config{
		AppName: getEnv("APP_NAME", ""),
		Port:    getEnv("APP_PORT", "8080"),
		Env:     getEnv("APP_ENV", "development"),
	}

	if cfg.AppName == "" {
		return cfg, errors.New("APP_NAME is required")
	}

	if cfg.Env != "development" && cfg.Env != "production" {
		return cfg, errors.New("APP_ENV must be development or production")
	}

	return cfg, nil
}
