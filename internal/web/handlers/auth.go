package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rcarvalho-pb/checkin-go/internal/config"
	"github.com/rcarvalho-pb/checkin-go/internal/participant"
	"github.com/rcarvalho-pb/checkin-go/internal/security"
)

func Login(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// req := struct {
		// 	Email    string `json:"email"`
		// 	Password string `json:"password"`
		// }{}
		if err := r.ParseForm(); err != nil {
			http.Error(w, "invalid login or password", http.StatusBadRequest)
			return
		}
		email := r.FormValue("email")
		password := r.FormValue("password")
		// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// 	http.Error(w, "invalid login or password", http.StatusBadRequest)
		// 	return
		// }
		p, err := app.ParticipantRepository.FindByEmail(email)
		// p, err := app.ParticipantRepository.FindByEmail(req.Email)
		if err != nil {
			http.Error(w, "invalid login or password", http.StatusBadRequest)
			return
		}
		if err := security.ValidatePassword(password, p.Password); err != nil {
			// if err := security.ValidatePassword(req.Password, p.Password); err != nil {
			http.Error(w, "invalid login or password", http.StatusBadRequest)
			return
		}
		token, err := app.AuthHandler.NewJwtToken(p)
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

func SignUp(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")
		hashedPassword, err := security.Hash(password)
		if err != nil {
			http.Error(w, "error encrypting password: "+err.Error(), http.StatusInternalServerError)
			return
		}
		p := &participant.Participant{
			Name:      name,
			Email:     email,
			Password:  string(hashedPassword),
			Active:    true,
			Role:      2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := app.ParticipantRepository.Create(p); err != nil {
			http.Error(w, "error creating participant:"+err.Error(), http.StatusInternalServerError)
			return
		}
		app.InfoLog.Println("user created successfuly")
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
