package routes

import (
	"net/http"

	"github.com/rcarvalho-pb/checkin-go/internal/auth"
	"github.com/rcarvalho-pb/checkin-go/internal/config"
)

func GetAuthRoutes(app ...*config.App) []Route {
	return []Route{
		{
			URI:            "/login",
			Method:         http.MethodPost,
			Function:       auth.Login(app[0]),
			Authentication: false,
		},
	}
}
