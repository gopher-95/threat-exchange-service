package handler

import (
	"errors"
	"net/http"

	"github.com/gopher-95/threat-exchange-service/internal/threat"
)

func handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, threat.ErrNotFound):
		writeErrorJSON(w, http.StatusNotFound, "Threat not found", err.Error())
		return
	case errors.Is(err, threat.ErrConflict):
		writeErrorJSON(w, http.StatusConflict, "Threat already exists", err.Error())
		return
	case errors.Is(err, threat.ErrInvalidMD5), errors.Is(err, threat.ErrInvalidFields):
		writeErrorJSON(w, http.StatusBadRequest, "Invalid request", err.Error())
		return
	default:
		writeErrorJSON(w, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}
}
