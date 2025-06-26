package storage

import (
	_ "github.com/lib/pq"
)

type Migrator interface {
	RunMigrationsUp()
}

func RunMigrationsUp(m Migrator) {
	m.RunMigrationsUp()
}

