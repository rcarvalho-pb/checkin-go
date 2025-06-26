package templates

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	Home().Render(r.Context(), w)
}
