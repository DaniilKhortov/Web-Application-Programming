package handler

import (
	"html/template"
	"log/slog"
	"net/http"
	"queueapp/internal/service"
	"queueapp/internal/storage"
	"strconv"
)

type QueueHandler struct {
	svc    *service.QueueService
	logger *slog.Logger
}

func NewQueueHandler(svc *service.QueueService, logger *slog.Logger) *QueueHandler {
	return &QueueHandler{svc: svc, logger: logger}
}

func (h *QueueHandler) Home(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Request", "method", r.Method, "path", r.URL.Path)

	tmpl, err := template.ParseFiles("web/templates/queue.html")
	if err != nil {
		h.serverError(w, err)
		return
	}

	data := struct {
		Clients []storage.QueueItem
	}{
		Clients: h.svc.GetAll(),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		h.serverError(w, err)
	}
}

func (h *QueueHandler) Create(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl, err := template.ParseFiles("web/templates/create.html")
		if err != nil {
			h.serverError(w, err)
			return
		}
		tmpl.Execute(w, nil)

	case http.MethodPost:
		name := r.FormValue("name")
		if name == "" {
			h.clientError(w, http.StatusBadRequest)
			return
		}
		h.svc.AddClient(name)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		h.clientError(w, http.StatusMethodNotAllowed)
	}
}

func (h *QueueHandler) Edit(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl, _ := template.ParseFiles("web/templates/edit.html")
		tmpl.Execute(w, r.URL.Query().Get("id"))
	case http.MethodPost:
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			h.clientError(w, http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")
		err = h.svc.EditClient(id, name)
		if err != nil {
			h.serverError(w, err)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (h *QueueHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		h.clientError(w, http.StatusBadRequest)
		return
	}
	err = h.svc.RemoveClient(id)
	if err != nil {
		h.serverError(w, err)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// --- centralized error handlers ---
func (h *QueueHandler) serverError(w http.ResponseWriter, err error) {
	h.logger.Error("server error", "error", err)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func (h *QueueHandler) clientError(w http.ResponseWriter, status int) {
	h.logger.Warn("client error", "status", status)
	http.Error(w, http.StatusText(status), status)
}
