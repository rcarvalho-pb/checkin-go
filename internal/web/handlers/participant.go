package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rcarvalho-pb/checkin-go/internal/auth"
	"github.com/rcarvalho-pb/checkin-go/internal/config"
	participant_role "github.com/rcarvalho-pb/checkin-go/internal/participant/roles"
)

func FindParticipatById(app ...*config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Here")
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "invalid user id", http.StatusBadRequest)
			return
		}
		app[0].InfoLog.Println("find by id:", id)
		p, err := app[0].ParticipantRepository.FindByID(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("error finding user: [%d]", id), http.StatusNotFound)
			return
		}
		response := struct {
			Status string `json:"status"`
			Data   any    `json:"data"`
		}{
			Status: "ok",
			Data:   p,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(response)
	}
}

func Teste(app ...*config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(auth.UserIDKey)
		name := r.Context().Value(auth.UserIDKey)
		email := r.Context().Value(auth.UserEmailKey)
		role := r.Context().Value(auth.UserRoleKey)
		response := struct {
			ID    int                   `json:"id"`
			Name  string                `json:"name"`
			Email string                `json:"email"`
			Role  participant_role.Role `json:"role"`
		}{
			ID:    id.(int),
			Name:  name.(string),
			Email: email.(string),
			Role:  role.(participant_role.Role),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
