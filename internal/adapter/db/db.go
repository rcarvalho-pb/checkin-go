package db

import (
	"log"

	"github.com/rcarvalho-pb/checkin-go/internal/adapter/postgres"
	"github.com/rcarvalho-pb/checkin-go/internal/participant"
)

type DB struct {
	participant.ParticipantRepository
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
		db := postgres.OpenDB(dsn)
		return &DB{db}
	default:
		return nil
	}
}
