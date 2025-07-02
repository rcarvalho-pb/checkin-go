package routes

import (
	"encoding/json"
	"net/http"
)

func GetHeathRoutes() []Route {
	return []Route{
		{
			URI:    "/health",
			Method: http.MethodGet,
			Function: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				resp := struct {
					Status string `json:"status"`
				}{Status: "ok"}
				json.NewEncoder(w).Encode(resp)
			},
		},
	}
}
