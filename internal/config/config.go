package config

import (
	"github.com/rcarvalho-pb/checkin-go/internal/event"
	"github.com/rcarvalho-pb/checkin-go/internal/participant"
)

type Migrator interface {
	RunMigrationsUp()
}

type App struct {
	EventHandler       *event.EventHandler
	ParticipantHandler *participant.ParticipantHandler
}

func (app *App) RunMigrationsUp(m Migrator) {
	m.RunMigrationsUp()
}
