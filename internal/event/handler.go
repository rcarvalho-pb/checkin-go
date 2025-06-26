package event

import (
	"net/http"
	"strconv"
	"time"
)

type EventHandler struct {
	es eventStorage
}

func NewEventHandler(e eventStorage) *EventHandler {
	return &EventHandler{
		e,
	}
}

func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "error parsing form", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	location := r.FormValue("location")
	startsAt, _ := time.Parse("2006-01-02T15:04", r.FormValue("starts_at"))
	endsAt, _ := time.Parse("2006-01-02T15:04", r.FormValue("ends_at"))

	event := Event{
		Name:     name,
		Location: location,
		StartsAt: startsAt,
		EndsAt:   endsAt,
	}
	id, err := h.es.Create(r.Context(), &event)
	if err != nil {
		http.Error(w, "error creating event", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(`<p style="color: green;">✅ Evento criado com sucesso! ID:` + strconv.Itoa(id) + ` </p>`))
}
