package storage

import (
	_ "github.com/lib/pq"
)

type Migrator interface {
	RunMigrationsUP()
}

