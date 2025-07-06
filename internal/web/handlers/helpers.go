package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rcarvalho-pb/checkin-go/internal/config"
)

func isAppPresent(app ...*config.App) bool {
	if len(app) < 1 {
		return false
	}
	return true
}

func JSON(w http.ResponseWriter, status string, statusCod int, data ...any) {
	resp := struct {
		Status string `json:"status"`
		Data   any    `json:"data,omitempty"`
	}{Status: status}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCod)
	if len(data) > 0 {
		resp.Data = data[0]
	}
	json.NewEncoder(w).Encode(resp)
}
