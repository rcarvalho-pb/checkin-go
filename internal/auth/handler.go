package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rcarvalho-pb/checkin-go/internal/config"
	"github.com/rcarvalho-pb/checkin-go/internal/participant"
	"github.com/rcarvalho-pb/checkin-go/internal/security"
)

func Login(app ...*config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(app) < 0 {
			http.Error(w, "app not available", http.StatusInternalServerError)
			return
		}
		if err := r.ParseForm(); err != nil {
			http.Error(w, "invalid login or password", http.StatusBadRequest)
			return
		}
		email := r.FormValue("email")
		password := r.FormValue("password")
		p, err := app[0].ParticipantRepository.FindByEmail(email)
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

func SignUp(app ...*config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(app) < 0 {
			http.Error(w, "app not available", http.StatusInternalServerError)
			return
		}
		if err := r.ParseForm(); err != nil {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")
		p := &participant.Participant{
			Name:      name,
			Email:     email,
			Password:  password,
			Active:    true,
			Role:      2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := app[0].ParticipantRepository.Create(p); err != nil {
			http.Error(w, "error creating participant", http.StatusInternalServerError)
			return
		}
		response := struct {
			Status string `json:"status"`
		}{
			Status: "created",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}
