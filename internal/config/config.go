package config

import (
	"github.com/rcarvalho-pb/checkin-go/internal/handler"
	"github.com/rcarvalho-pb/checkin-go/internal/templates"
)

type Migrator interface {
	RunMigrationsUp()
}

type App struct {
	EventHandler       *handler.EventHandler
	ParticipantHandler *handler.ParticipantHandler
	TemplateHandler    *templates.TemplateHandler
}

func (app *App) RunMigrationsUp(m Migrator) {
	m.RunMigrationsUp()
}
