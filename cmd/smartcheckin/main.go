package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rcarvalho-pb/checkin-go/internal/config"
	"github.com/rcarvalho-pb/checkin-go/internal/db"
	"github.com/rcarvalho-pb/checkin-go/internal/globals"
	"github.com/rcarvalho-pb/checkin-go/internal/web"
)

const webPort = "8080"

func main() {
	config.StartApp()
	dbPool := db.GetDB(globals.DBType, globals.DSN)
	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	app := &config.App{
		ParticipantRepository: dbPool.ParticipantRepository,
		EventRepository:       dbPool.EventRepository,
		InfoLog:               infoLog,
		ErrorLog:              errorLog,
	}
	r := web.StartRouter(app)
	app.InfoLog.Printf("Server started on port: %s", webPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", webPort), r); err != nil {
		app.ErrorLog.Fatal(err)
	}
}
