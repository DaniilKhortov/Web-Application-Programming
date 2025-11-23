package handler

import (
	"html/template"
	"log/slog"
	"net/http"
	"sync"
)

// QueueHandler — структура з логером і сховищем
type QueueHandler struct {
	logger  *slog.Logger
	mu      sync.Mutex
	clients []string
}

// Конструктор
func NewQueueHandler(logger *slog.Logger) *QueueHandler {
	return &QueueHandler{
		logger: logger,
		clients: []string{
			"Ivan Petrenko",
			"Maria Kondratenko",
			"Oleh Stupa",
			"Volodymyr Kononenko",
		},
	}
}

// ShowQueue — Read-операція
func (h *QueueHandler) ShowQueue(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling request",
		"method", r.Method,
		"path", r.URL.Path,
	)

	tmpl, err := template.ParseFiles("web/templates/queue.html")
	if err != nil {
		h.logger.Error("Error parsing template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	data := struct {
		Title   string
		Clients []string
	}{
		Title:   "E-Queue",
		Clients: h.clients,
	}

	if err := tmpl.Execute(w, data); err != nil {
		h.logger.Error("Error executing template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// AddClient — Create-операція
func (h *QueueHandler) AddClient(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling request",
		"method", r.Method,
		"path", r.URL.Path,
	)

	switch r.Method {
	case http.MethodGet:
		tmpl, err := template.ParseFiles("web/templates/add.html")
		if err != nil {
			h.logger.Error("Error parsing add.html", "error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)

	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			h.logger.Error("Error parsing form", "error", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		name := r.PostFormValue("name")
		if name == "" {
			http.Error(w, "Name cannot be empty", http.StatusBadRequest)
			return
		}

		h.mu.Lock()
		h.clients = append(h.clients, name)
		h.mu.Unlock()

		h.logger.Info("Client added", "name", name, "total", len(h.clients))

		// PRG-патерн: редирект після POST
		http.Redirect(w, r, "/", http.StatusSeeOther)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
