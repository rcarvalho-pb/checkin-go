package router

import (
	"net/http"

	"github.com/rcarvalho-pb/checkin-go/internal/router/routes"
)

func GetRouter() *http.ServeMux {
	mux := http.NewServeMux()
	routes.Config(mux)
	return mux
}
