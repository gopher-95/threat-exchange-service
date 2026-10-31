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

type ErrorResponse struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

func NewHandler(serv Service) *Handler {
	return &Handler{serv: serv}
}

func (h *Handler) GetInfoHandle(w http.ResponseWriter, r *http.Request) {
	md5 := chi.URLParam(r, "md5")

	ctx := r.Context()

	_, err := h.serv.GetInfoService(ctx, md5)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: "такой угрозы не найден",
			Details: "нет угрозы с таким md5",
		})
	}

}
