package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// Helper functions to reduce duplication
func ParseIDFromRequest(w http.ResponseWriter, r *http.Request) (int, bool) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		SendError(w, http.StatusBadRequest, "Invalid user ID")
		return 0, false
	}
	return id, true
}

func SendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func SendError(w http.ResponseWriter, statusCode int, message string) {
	SendJSONResponse(w, statusCode, ErrorResponse{Error: message})
}

func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dest interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		SendError(w, http.StatusBadRequest, "Invalid request payload")
		return false
	}
	return true
}
