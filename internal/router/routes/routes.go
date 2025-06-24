package routes

import (
	"fmt"
	"net/http"
)

type Route struct {
	Uri            string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	Authentication bool
	Admin          bool
}

func Config(r *http.ServeMux) *http.ServeMux {
	var routes []Route
	routes = append(routes, EventRoutes...)

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
