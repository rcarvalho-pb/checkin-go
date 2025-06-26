package templates

import "net/http"

type TemplateHandler struct{}

func (t *TemplateHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	t.home().Render(r.Context(), w)
}
