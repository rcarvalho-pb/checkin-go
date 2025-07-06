package routes

import (
	"encoding/json"
	"net/http"

	"github.com/rcarvalho-pb/checkin-go/internal/config"
)

func getHeathRoutes() []Route {
	return []Route{
		{
			URI:    "/health",
			Method: http.MethodGet,
			Function: func(app *config.App) http.HandlerFunc {
				_ = app
				return func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					resp := struct {
						Status string `json:"status"`
					}{Status: "ok"}
					json.NewEncoder(w).Encode(resp)
				}
			},
			Authentication: false,
		},
	}
}
