package handlers

import (
	"net/http"

	"github.com/rcarvalho-pb/checkin-go/internal/templates"
)

func CreateEventPage(w http.ResponseWriter, r *http.Request) {
	templates.EventForm().Render(r.Context(), w)
}

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {

}
