package config

import "fmt"

//Структура Config зберігає налаштування черги
type Config struct {
	MaxSize int
}

//Функція Load завантажує конфігурацію програми
//Наразі параметри черги задаються вручну
func Load() *Config {
	cfg := &Config{MaxSize: 5}
	fmt.Println("Configuration finished!")
	return cfg
}
