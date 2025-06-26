package postgres

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var dbTimeout = 10 * time.Second
var DSN string

var DSN string

type DB struct {
	*sqlx.DB
}

func GetDB(dsn string) *DB {
	DSN = dsn
	db := connectToDB()
	if db == nil {
		log.Fatal("couldn`t connect to DB")
	}
	return &DB{
		db,
	}
}

func connectToDB() *sqlx.DB {
	count := 0
	for count < 10 {
		db, err := openDB()
		if err == nil {
			return db
		}
		count++
		time.Sleep(1 * time.Second)
	}
	return nil
}

func openDB() (*sqlx.DB, error) {
	conn, err := sqlx.Open("postgres", DSN)
	if err != nil {
		return nil, err
	}
	if err = conn.Ping(); err != nil {
		return nil, err
	}
	return conn, nil
}

func (db *DB) RunMigrationsUp() {
	rootDir, err := os.Executable()
	if err != nil {
		log.Fatalf("error getting cwd: %v", err)
	}
	migrationsPath := filepath.Join("file://", filepath.Dir(filepath.Dir(rootDir)), "migrations")
	m, err := migrate.New(
		migrationsPath,
		DSN,
	)
	if err != nil {
		log.Fatalf("error creating migrator: %v", err)
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("error running migrations: %v", err)
	}
	fmt.Println("migrations successfully applied")
}
