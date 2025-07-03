package auth

import (
	"encoding/json"
	"net/http"

	"github.com/rcarvalho-pb/checkin-go/internal/config"
	"github.com/rcarvalho-pb/checkin-go/internal/security"
)

func Login(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "invalid login or password", http.StatusBadRequest)
			return
		}
		email := r.FormValue("email")
		password := r.FormValue("password")
		p, err := app.ParticipantRepository.FindByEmail(email)
		if err != nil {
			http.Error(w, "invalid login or password", http.StatusBadRequest)
			return
		}
		if err := security.ValidatePassword(password, p.Password); err != nil {
			http.Error(w, "invalid login or password", http.StatusBadRequest)
			return
		}
		token, err := NewJwtToken(p)
		if err != nil {
			http.Error(w, "error generating jwt token", http.StatusInternalServerError)
			return
		}
		response := struct {
			Token string `json:"token"`
		}{
			Token: token,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
