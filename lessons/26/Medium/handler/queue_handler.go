package handler

import (
	"encoding/json"
	"net/http"

	"queueapp/service"
)

type QueueHandler struct {
	Service service.QueueService
}

func NewQueueHandler(s service.QueueService) *QueueHandler {
	return &QueueHandler{Service: s}
}

func (h *QueueHandler) GetClientsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	clients := h.Service.GetAllClients()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clients)
}
