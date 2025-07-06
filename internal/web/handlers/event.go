package handlers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/rcarvalho-pb/checkin-go/internal/auth"
	"github.com/rcarvalho-pb/checkin-go/internal/config"
	"github.com/rcarvalho-pb/checkin-go/internal/event"
)

func CreateEvent(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "invalid inputs", http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")
		location := r.FormValue("location")
		ownerID, ok := r.Context().Value(auth.UserIDKey).(int)
		if !ok {
			http.Error(w, "invalid user", http.StatusNotAcceptable)
			return
		}
		startsAt, err := time.Parse("2006-01-02", r.FormValue("starts_at"))
		if err != nil {
			http.Error(w, "invalid date format", http.StatusBadRequest)
			return
		}
		endsAt, err := time.Parse("2006-01-02", r.FormValue("ends_at"))
		if err != nil {
			http.Error(w, "invalid date format", http.StatusBadRequest)
			return
		}
		if startsAt.After(endsAt) {
			http.Error(w, "cannot start after ends", http.StatusBadRequest)
			return
		}
		latitude, err := strconv.ParseFloat(r.FormValue("latitude"), 64)
		if err != nil {
			http.Error(w, "error converting latitude and longitude", http.StatusBadRequest)
			return
		}
		longitude, err := strconv.ParseFloat(r.FormValue("longitude"), 64)
		if err != nil {
			http.Error(w, "error converting latitude and longitude", http.StatusBadRequest)
			return
		}
		e := &event.Event{
			Name:      name,
			Location:  location,
			StartsAt:  startsAt,
			EndsAt:    endsAt,
			OwnerID:   ownerID,
			Latitude:  latitude,
			Longitude: longitude,
		}
		if err := app.EventRepository.Create(e); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		JSON(w, "ok", http.StatusOK)
	}
}

func FindEvents(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := app.EventRepository.FindAll()
		if err != nil {
			http.Error(w, "error finding events", http.StatusInternalServerError)
			return
		}
		JSON(w, "ok", http.StatusOK, events)
	}
}

func Checkin(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		participantID := r.Context().Value(auth.UserIDKey).(int)
		if err := r.ParseForm(); err != nil {
			http.Error(w, "invalid inputs", http.StatusBadRequest)
			return
		}
		eventID, err := strconv.Atoi(r.FormValue("event_id"))
		if err != nil {
			http.Error(w, "invalid inputs", http.StatusBadRequest)
			return
		}
		latitude, err := strconv.ParseFloat(r.FormValue("latitude"), 64)
		if err != nil {
			http.Error(w, "error converting latitude and longitude", http.StatusBadRequest)
			return
		}
		longitude, err := strconv.ParseFloat(r.FormValue("longitude"), 64)
		if err != nil {
			http.Error(w, "error converting latitude and longitude", http.StatusBadRequest)
			return
		}
		e, err := app.EventRepository.FindByID(eventID)
		if err != nil {
			http.Error(w, fmt.Sprintf("event [%d]: not found", eventID), http.StatusInternalServerError)
			return
		}
		dist := haversine(e.Latitude, e.Longitude, latitude, longitude)
		if dist > 100.0 {
			http.Error(w, "can't checkin in another plance than the event", http.StatusBadRequest)
			return
		}
		if err := app.EventRepository.Checkin(eventID, participantID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		JSON(w, "created", http.StatusCreated, struct {
			Distance float64 `json:"distance"`
		}{Distance: dist})
	}
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000 // raio da Terra em metros
	dLat := (lat2 - lat1) * math.Pi / 180.0
	dLon := (lon2 - lon1) * math.Pi / 180.0

	lat1 = lat1 * math.Pi / 180.0
	lat2 = lat2 * math.Pi / 180.0

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c // distância em metros
}
