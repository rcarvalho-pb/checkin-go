package routes

import (
	"fmt"
	"net/http"

	"github.com/rcarvalho-pb/checkin-go/internal/config"
)

type Route struct {
	URI            string
	Method         string
	Function       func(...*config.App) http.HandlerFunc
	Authentication bool
}

func ConfigRoutes(mux *http.ServeMux, app *config.App) *http.ServeMux {
	routes := []Route{}
	routes = append(routes, GetHeathRoutes()...)
	routes = append(routes, GetAuthRoutes(app)...)
	for _, r := range routes {
		if r.Authentication {

		} else {
			mux.HandleFunc(fmt.Sprintf("%s %s", r.Method, r.URI), r.Function(app))
		}
	}
	return mux
}
