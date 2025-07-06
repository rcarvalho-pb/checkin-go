package routes

import (
	"net/http"

	"github.com/rcarvalho-pb/checkin-go/internal/web/handlers"
)

const EVENT_RESOURCE = "/events"

func getEventRoutes() []Route {
	return []Route{
		{
			URI:            EVENT_RESOURCE,
			Method:         http.MethodPost,
			Function:       handlers.CreateEvent,
			Authentication: true,
		},
		{
			URI:            EVENT_RESOURCE,
			Method:         http.MethodGet,
			Function:       handlers.FindEvents,
			Authentication: false,
		},
		{
			URI:            EVENT_RESOURCE + "/checkin",
			Method:         http.MethodPost,
			Function:       handlers.Checkin,
			Authentication: true,
		},
	}
}
