package routes

import (
	"net/http"

	"github.com/rcarvalho-pb/checkin-go/internal/web/handlers"
)

func getAuthRoutes() []Route {
	return []Route{
		{
			URI:            "/login",
			Method:         http.MethodPost,
			Function:       handlers.Login,
			Authentication: false,
		},
		{
			URI:            "/signup",
			Method:         http.MethodPost,
			Function:       handlers.SignUp,
			Authentication: false,
		},
	}
}
