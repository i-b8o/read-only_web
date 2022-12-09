package v1

import (
	"encoding/json"
	"net/http"
)

type statusResponse struct {
	Message string `json:"status"`
}

func newErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(statusResponse{Message: message})
}

func newStatusResponse(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(statusResponse{Message: message})
}
