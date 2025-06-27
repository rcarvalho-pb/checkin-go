package routes

import (
	"net/http"

	"github.com/rcarvalho-pb/checkin-go/internal/config"
)

func getEventRoutes(app *config.App) []Route {
	return []Route{
		{
			Uri:            "/",
			Method:         http.MethodGet,
			Function:       app.TemplateHandler.HomeHandler,
			Authentication: false,
			Admin:          false,
		},
		{
			Uri:            "/events",
			Method:         http.MethodGet,
			Function:       app.TemplateHandler.CreateEventPage,
			Authentication: false,
			Admin:          false,
		},
	}
}
