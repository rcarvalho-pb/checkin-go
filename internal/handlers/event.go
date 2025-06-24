package handlers

import (
	"net/http"

	"github.com/rcarvalho-pb/checkin-go/internal/templates"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	templates.EventForm().Render(r.Context(), w)
}
