package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gopher-95/threat-exchange-service/internal/threat"
)

// Service определяет контракт для бизнес-логики, используемой хендлерами.
type Service interface {
	GetInfoService(ctx context.Context, md5 string) (*threat.Threat, error)
	CreateInfoService(ctx context.Context, t *threat.ThreatCreateRequest) (*threat.Threat, error)
	UpdateInfoService(ctx context.Context, md5 string, t *threat.ThreatUpdateRequest) (*threat.Threat, error)
}

// Handler обрабатывает HTTP-запросы, связанные с угрозами.
type Handler struct {
	serv Service
}

// NewHandler создает новый экземпляр Handler с указанным сервисом.
func NewHandler(serv Service) *Handler {
	return &Handler{serv: serv}
}

// GetInfoHandle обрабатывает GET-запрос для получения информации об угрозе по MD5.
func (h *Handler) GetInfoHandle(w http.ResponseWriter, r *http.Request) {
	md5 := chi.URLParam(r, "md5")

	t, err := h.serv.GetInfoService(r.Context(), md5)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, t)
}

// CreateInfoHandle обрабатывает POST-запрос для создания новой угрозы.
func (h *Handler) CreateInfoHandle(w http.ResponseWriter, r *http.Request) {
	var req threat.ThreatCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorJSON(w, 400, "Invalid JSON", err.Error())
		return
	}

	t, err := h.serv.CreateInfoService(r.Context(), &req)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, t)
}

// UpdateInfoHandle обрабатывает PUT-запрос для обновления существующей угрозы.
func (h *Handler) UpdateInfoHandle(w http.ResponseWriter, r *http.Request) {
	var req threat.ThreatUpdateRequest

	md5 := chi.URLParam(r, "md5")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorJSON(w, http.StatusBadRequest, "Invalid JSON", err.Error())
		return
	}

	t, err := h.serv.UpdateInfoService(r.Context(), md5, &req)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, t)

}
