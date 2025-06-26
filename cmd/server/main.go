package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rcarvalho-pb/checkin-go/internal/adapter/db/postgres"
	"github.com/rcarvalho-pb/checkin-go/internal/router"
	"github.com/rcarvalho-pb/checkin-go/internal/storage"
)

const webPort = 8080

var migrator storage.Migrator

func main() {

	migrator = &postgres.Migrator{}
	migrator.RunMigrationsUP()
	r := router.GetRouter()
	fmt.Printf("Server started on port: %d\n", webPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", webPort), r))
}
