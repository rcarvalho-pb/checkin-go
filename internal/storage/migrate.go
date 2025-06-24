package storage

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/lib/pq"
)

func RunMigrationsUP(dsn string) {
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
	fmt.Println("migrations successfully applied")
}
