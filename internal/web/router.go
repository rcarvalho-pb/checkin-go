package web

import (
	"net/http"

	"github.com/rcarvalho-pb/checkin-go/internal/config"
	"github.com/rcarvalho-pb/checkin-go/internal/web/routes"
)

func StartRouter(app *config.App) *http.ServeMux {
	mux := http.NewServeMux()
	routes.ConfigRoutes(mux, app)
	return mux
}
