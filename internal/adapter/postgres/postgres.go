package postgres

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

const dbTimeout = 10 * time.Second

func connectToDB(dsn string) *sqlx.DB {
	counts := 0
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("db not yet read...")
		} else {
			log.Println("connected to the db")
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
