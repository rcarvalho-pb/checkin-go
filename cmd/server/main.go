package main

import (
	"log"
	"net/http"

	"github.com/rcarvalho-pb/checkin-go/internal/router"
)

func main() {
	r := router.GetRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
