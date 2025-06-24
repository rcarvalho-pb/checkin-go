package event

import "time"

type Event struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Location  string    `json:"location" db:"location"`
	StartsAt  time.Time `json:"starts_at" db:"starts_at"`
	EndsAt    time.Time `json:"ends_at" db:"ends_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
