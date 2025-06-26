package postgres

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var dbTimeout = 10 * time.Second
var DSN string

type DB struct {
	*sqlx.DB
}

func GetDB(dsn string) *DB {
	db := connectToDB(dsn)
	DSN = dsn
	if db == nil {
		log.Fatal("couldn`t connect to DB")
	}
	return &DB{
		db,
	}
}

func connectToDB(dsn string) *sqlx.DB {
	count := 0
	for count < 10 {
		db, err := openDB(dsn)
		if err == nil {
			return db
		}
		count++
		time.Sleep(1 * time.Second)
	}
	return nil
}

func openDB(dsn string) (*sqlx.DB, error) {
	conn, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = conn.Ping(); err != nil {
		return nil, err
	}
	return conn, nil
}
