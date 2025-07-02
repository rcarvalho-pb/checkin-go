package routes

import (
	"fmt"
	"net/http"
)

type Route struct {
	URI            string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	Authentication bool
}

func ConfigRoutes(mux *http.ServeMux) *http.ServeMux {
	routes := []Route{}
	routes = append(routes, GetHeathRoutes()...)
	for _, r := range routes {
		if r.Authentication {

		} else {
			mux.HandleFunc(fmt.Sprintf("%s %s", r.Method, r.URI), r.Function)
		}
	}
	return mux
}
