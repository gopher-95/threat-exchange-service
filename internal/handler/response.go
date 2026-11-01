package handler

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse представляет стандартный формат JSON-ответа при возникновении ошибки.
type ErrorResponse struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

// writeJSON отправляет клиенту JSON-ответ с указанным HTTP-статусом.
func writeJSON(w http.ResponseWriter, statuscode int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	json.NewEncoder(w).Encode(v)
}

// writeErrorJSON отправляет клиенту JSON-ответ с ошибкой и указанным HTTP-статусом.
func writeErrorJSON(w http.ResponseWriter, statuscode int, message, details string) {
	writeJSON(w, statuscode, ErrorResponse{Message: message, Details: details})
}
