package routes

import (
	"net/http"

	"github.com/rcarvalho-pb/checkin-go/internal/web/handlers"
)

func getParticipantsRoutes() []Route {
	return []Route{
		{
			URI:            "/participants/find/{id}",
			Method:         http.MethodGet,
			Function:       handlers.FindParticipatById,
			Authentication: true,
		},
	}
}
