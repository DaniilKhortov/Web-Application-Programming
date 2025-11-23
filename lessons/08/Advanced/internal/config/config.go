package config

import "fmt"

type Config struct {
	MaxSize int
}

func Load() *Config {
	cfg := &Config{MaxSize: 5}
	fmt.Println("Configuration finished!")
	return cfg
}
