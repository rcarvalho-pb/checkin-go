package routes

import (
	"net/http"

	"github.com/rcarvalho-pb/checkin-go/internal/handlers"
)

var EventRoutes = []Route{
	{
		Uri:            "/",
		Method:         http.MethodGet,
		Function:       handlers.CreateEvent,
		Authentication: false,
		Admin:          false,
	},
}
