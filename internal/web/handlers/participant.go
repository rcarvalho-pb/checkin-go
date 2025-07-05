package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rcarvalho-pb/checkin-go/internal/config"
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
