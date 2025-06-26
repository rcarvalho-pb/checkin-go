package router

import (
	"net/http"

	"github.com/rcarvalho-pb/checkin-go/internal/config"
	"github.com/rcarvalho-pb/checkin-go/internal/router/routes"
)

func GetRouter(app *config.App) *http.ServeMux {
	mux := http.NewServeMux()
	routes.Config(mux, app)
	return mux
}
