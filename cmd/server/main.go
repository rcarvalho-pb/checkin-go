package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rcarvalho-pb/checkin-go/internal/adapter/db/postgres"
<<<<<<< HEAD
=======
	"github.com/rcarvalho-pb/checkin-go/internal/config"
	"github.com/rcarvalho-pb/checkin-go/internal/event"
	"github.com/rcarvalho-pb/checkin-go/internal/participant"
>>>>>>> d26c2a609706da1ce09d0a10f63584de2cab9b15
	"github.com/rcarvalho-pb/checkin-go/internal/router"
	"github.com/rcarvalho-pb/checkin-go/internal/storage"
)

const webPort = 8080

var migrator storage.Migrator

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
	fmt.Printf("Server started on port: %d\n", webPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", webPort), r))
}
