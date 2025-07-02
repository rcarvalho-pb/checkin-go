package config

import (
	"github.com/rcarvalho-pb/checkin-go/internal/helper"
	"github.com/rcarvalho-pb/checkin-go/internal/participant"
)

var (
	Secret, DBType, DSN string
)

type App struct {
	participant.ParticipantRepository
}

type Migrator interface {
	RunMigrationsUp(string)
}

func (app *App) RunMigrationsUp(dsn string, m Migrator) {
	m.RunMigrationsUp(dsn)
}

func StartApp() {
	Secret = helper.GetEnvWithCallback("SECRET", "verysecret")
	DBType = helper.GetEnvWithCallback("DB_TYPE", "sqlite")
	DSN = helper.GetEnvWithCallback("DSN", "../db-data/sqlite/sqlite.db")
}
