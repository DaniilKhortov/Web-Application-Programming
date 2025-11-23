package service

import (
	"errors"
	"regexp"
	"strings"
)

// ValidateClientName перевіряє, чи ім’я клієнта коректне:
// - не порожнє
// - довжина від 2 до 50 символів
// - не містить цифр або спецсимволів
func ValidateClientName(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.New("name cannot be empty")
	}
	if len(name) < 2 || len(name) > 50 {
		return errors.New("name length must be in between 2 and 50")
	}
	match, _ := regexp.MatchString(`^[A-Za-zА-Яа-яІіЇїЄє' -]+$`, name)
	if !match {
		return errors.New("name contain invalid symbols")
	}
	return nil
}
