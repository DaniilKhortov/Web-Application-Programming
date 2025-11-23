package service

import (
	"strings"
	"testing"
)

// Table-driven test для ValidateClientName
func TestValidateClientName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantErr  bool
		wantText string
	}{
		{"Valid short name", "Olha", false, ""},
		{"Valid long name", "Andriy-/Petro /Ivanovych", false, ""},
		{"Empty name", "", true, "name cannot be empty"},
		{"Too short", "A", true, "name length"},
		{"Too long", "fhdjhrkbdnbjdrnkbndrknb", true, "name length"},
		{"Contains digits", "Марія123", true, "invalid symbols"},
		{"Contains special chars", "Оля!", true, "invalid symbols"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateClientName(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("очікувалась помилка, але отримано nil")
				}
				if !contains(err.Error(), tt.wantText) {
					t.Errorf("очікувана помилка містить %q, отримано %q", tt.wantText, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("очікувалось без помилки, отримано: %v", err)
				}
			}
		})
	}
}

// допоміжна функція для пошуку фрагмента тексту
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
