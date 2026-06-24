package api

import (
	"encoding/json"
	"net/http"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {

	json.NewEncoder(w).Encode(
		map[string]string{
			"status": "healthy",
		},
	)
}
