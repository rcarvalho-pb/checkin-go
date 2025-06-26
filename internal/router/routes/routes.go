package routes

import (
	"fmt"
	"net/http"

	"github.com/rcarvalho-pb/checkin-go/internal/config"
)

type Route struct {
	Uri            string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	Authentication bool
	Admin          bool
}

func Config(r *http.ServeMux, app *config.App) *http.ServeMux {
	fmt.Printf("%+v\n", app)
	var routes []Route
	routes = append(routes, getEventRoutes(app)...)

	for _, route := range routes {
		if route.Authentication {
			if route.Admin {

			} else {

			}
		} else {
			r.HandleFunc(fmt.Sprintf("%s %s", route.Method, route.Uri), route.Function)
		}
	}
	return r
}
