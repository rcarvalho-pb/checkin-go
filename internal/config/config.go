package config

import (
	"log"

	"github.com/rcarvalho-pb/checkin-go/internal/auth"
	"github.com/rcarvalho-pb/checkin-go/internal/event"
	"github.com/rcarvalho-pb/checkin-go/internal/globals"
	"github.com/rcarvalho-pb/checkin-go/internal/helper"
	"github.com/rcarvalho-pb/checkin-go/internal/participant"
)

var ()

type App struct {
	participant.ParticipantRepository
	event.EventRepository
	AuthHandler *auth.AuthHandler
	InfoLog     *log.Logger
	ErrorLog    *log.Logger
}

type Migrator interface {
	RunMigrationsUp(string)
}

func (app *App) RunMigrationsUp(dsn string, m Migrator) {
	m.RunMigrationsUp(dsn)
}

func StartApp() {
	globals.Secret = helper.GetEnvWithCallback("SECRET", "verysecret")
	globals.DBType = helper.GetEnvWithCallback("DB_TYPE", "sqlite")
	globals.DSN = helper.GetEnvWithCallback("DSN", "../db-data/sqlite/sqlite.db")
}
