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
	routes = append(routes, getHeathRoutes()...)
	routes = append(routes, getAuthRoutes()...)
	routes = append(routes, getParticipantsRoutes()...)
	routes = append(routes, getEventRoutes()...)
	for _, r := range routes {
		if r.Authentication {
			mux.HandleFunc(fmt.Sprintf("%s %s", r.Method, r.URI), app.AuthHandler.LoggerMiddleware(app.AuthHandler.AuthMiddleware(r.Function(app))))
		} else {
			mux.HandleFunc(fmt.Sprintf("%s %s", r.Method, r.URI), app.AuthHandler.LoggerMiddleware(r.Function(app)))
		}
	}
	return mux
}
