package templates

import (
	"net/http"
)

type TemplateHandler struct{}

func NewTemplateHandler() *TemplateHandler {
	return &TemplateHandler{}
}

func (t *TemplateHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	t.home().Render(r.Context(), w)
}

func (t *TemplateHandler) CreateEventPage(w http.ResponseWriter, r *http.Request) {
	eventForm().Render(r.Context(), w)
}
