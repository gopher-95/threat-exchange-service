package handler

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

func writeJSON(w http.ResponseWriter, statuscode int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	json.NewEncoder(w).Encode(v)
}

func writeErrorJSON(w http.ResponseWriter, statuscode int, message, details string) {
	writeJSON(w, statuscode, ErrorResponse{Message: message, Details: details})
}
