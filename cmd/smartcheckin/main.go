package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rcarvalho-pb/checkin-go/internal/adapter/db"
	"github.com/rcarvalho-pb/checkin-go/internal/config"
	"github.com/rcarvalho-pb/checkin-go/internal/web"
)

const webPort = "8080"

func main() {
	config.StartApp()
	dbPool := db.GetDB(config.DBType, config.DSN)
	infoLog := log.New(os.Stdout, "INFO: ")
	app := *&config.App{
		ParticipantRepository: dbPool,
	}
	r := web.StartRouter()
	log.Printf("Server started on port: %s", webPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", webPort), r); err != nil {
		log.Fatal(err)
	}
}
