package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gopher-95/threat-exchange-service/internal/threat"
)

type Service interface {
	GetInfoService(ctx context.Context, md5 string) (*threat.Threat, error)
	CreateInfoService(ctx context.Context, t *threat.ThreatCreateRequest) (*threat.Threat, error)
	UpdateInfoService(ctx context.Context, md5 string, t *threat.ThreatUpdateRequest) (*threat.Threat, error)
}

type Handler struct {
	serv Service
}

func NewHandler(serv Service) *Handler {
	return &Handler{serv: serv}
}

func (h *Handler) GetInfoHandle(w http.ResponseWriter, r *http.Request) {
	md5 := chi.URLParam(r, "md5")

	t, err := h.serv.GetInfoService(r.Context(), md5)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, t)
}

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
