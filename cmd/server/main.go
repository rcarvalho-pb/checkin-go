package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rcarvalho-pb/checkin-go/internal/adapter/db/postgres"
	"github.com/rcarvalho-pb/checkin-go/internal/config"
	"github.com/rcarvalho-pb/checkin-go/internal/event"
	"github.com/rcarvalho-pb/checkin-go/internal/participant"
	"github.com/rcarvalho-pb/checkin-go/internal/router"
)

func main() {
	dsn := os.Getenv("DSN")
	db := postgres.GetDB(dsn)
	eventSt := postgres.NewEventStorage(db)
	participantSt := postgres.NewParticipantStorage(db)
	app := config.App{
		EventHandler:       event.NewEventHandler(eventSt),
		ParticipantHandler: participant.NewParticipantHandler(participantSt),
	}
	r := router.GetRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
