package web

import (
	"net/http"

	"github.com/rcarvalho-pb/checkin-go/internal/web/routes"
)

func StartRouter() *http.ServeMux {
	mux := http.NewServeMux()
	routes.ConfigRoutes(mux)
	return mux
}
