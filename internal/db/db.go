package db

import (
	"log"

	"github.com/rcarvalho-pb/checkin-go/internal/adapter/postgres"
	"github.com/rcarvalho-pb/checkin-go/internal/event"
	"github.com/rcarvalho-pb/checkin-go/internal/participant"
)

type DB struct {
	participant.ParticipantRepository
	event.EventRepository
}

func GetDB(dbType, dsn string) *DB {
	conn := selectDB(dbType, dsn)
	if conn == nil {
		log.Panic("can't connect to database")
	}
	return conn
}

func selectDB(dbType, dsn string) *DB {
	switch dbType {
	case "postgres":
		dbParticipant, dbEvent := postgres.GetRepositories(dsn)
		return &DB{dbParticipant, dbEvent}
	default:
		return nil
	}
}
