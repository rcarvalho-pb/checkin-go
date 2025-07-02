package db

import (
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
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

func (db *DB) RunMigrationUp(dsn string) {
	rootDir, err := os.Executable()
	if err != nil {
		log.Fatalf("error getting cwd: %v", err)
	}
	migrationsPath := filepath.Join("file://", filepath.Dir(filepath.Dir(rootDir)), "migrations")
	m, err := migrate.New(
		migrationsPath,
		dsn,
	)
	if err != nil {
		log.Fatalf("error creating migrator: %v", err)
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("error running migrations: %v", err)
	}
	log.Println("migrations successfully applied")
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
