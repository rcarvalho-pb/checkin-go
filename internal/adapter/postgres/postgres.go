package postgres

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

const dbTimeout = 10 * time.Second

func GetRepositories(dsn string) (*PostgresParticipant, *PostgresEvent) {
	db := connectToDB(dsn)
	return &PostgresParticipant{db}, &PostgresEvent{db}
}

func connectToDB(dsn string) *sqlx.DB {
	counts := 0
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("db not yet read...")
		} else {
			log.Println("connected to the db")
			runMigrationUp(dsn)
			return connection
		}
		if counts > 10 {
			return nil
		}
		log.Println("Backing off for 1 sec...")
		time.Sleep(1 * time.Second)
		counts++
	}
}

func openDB(dsn string) (*sqlx.DB, error) {
	conn, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func runMigrationUp(dsn string) {
	execPath, err := os.Executable()
	if err != nil {
		log.Fatalf("error getting cwd: %v", err)
	}
	migrationsDir := filepath.Join(filepath.Dir(filepath.Dir(execPath)), "migrations")
	migrationsPath := "file://" + migrationsDir
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
