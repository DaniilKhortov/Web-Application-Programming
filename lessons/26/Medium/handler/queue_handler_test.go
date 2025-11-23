package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockQueueService struct {
	mockData []string
}

func (m *mockQueueService) GetAllClients() []string {
	return m.mockData
}

func TestGetClientsHandler(t *testing.T) {
	mockService := &mockQueueService{
		mockData: []string{"Olha", "Ivan", "Maria"},
	}

	handler := NewQueueHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/clients", nil)
	rr := httptest.NewRecorder()

	handler.GetClientsHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Expected status 200, received %d", rr.Code)
	}

	if !strings.Contains(rr.Header().Get("Content-Type"), "application/json") {
		t.Errorf("expected Content-Type to contain application/json, got %q", rr.Header().Get("Content-Type"))
	}

	body := rr.Body.String()

	var result []string
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		t.Fatalf("couldn`t parse body: %v", err)
	}

	expected := mockService.mockData
	if len(result) != len(expected) {
		t.Errorf("expected %d elements, received %d", len(expected), len(result))
	}

	for _, name := range expected {
		if !contains(body, name) {
			t.Errorf("expected client %q in response, but wasn`t there", name)
		}
	}
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
