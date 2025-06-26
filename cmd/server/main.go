package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rcarvalho-pb/checkin-go/internal/adapter/db/postgres"
	"github.com/rcarvalho-pb/checkin-go/internal/config"
	"github.com/rcarvalho-pb/checkin-go/internal/handler"
	"github.com/rcarvalho-pb/checkin-go/internal/router"
	"github.com/rcarvalho-pb/checkin-go/internal/templates"
)

const webPort = 8080

func main() {
	dsn := os.Getenv("DSN")
	db := postgres.GetDB(dsn)
	eventSt := postgres.NewEventStorage(db)
	participantSt := postgres.NewParticipantStorage(db)
	app := &config.App{
		EventHandler:       handler.NewEventHandler(eventSt),
		ParticipantHandler: handler.NewParticipantHandler(participantSt),
		TemplateHandler:    &templates.TemplateHandler{},
	}
	app.RunMigrationsUp(db)
	r := router.GetRouter(app)
	fmt.Printf("Server started on port: %d\n", webPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", webPort), r))
}
