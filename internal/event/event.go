package event

import "time"

type Event struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Location  string    `json:"location" db:"location"`
	StartsAt  time.Time `json:"starts_at" db:"starts_at"`
	EndsAt    time.Time `json:"ends_at" db:"ends_at"`
	OwnerID   int       `json:"owner_id" db:"owner_id"`
	Latitude  float64   `json:"latitude" db:"latitude"`
	Longitude float64   `json:"longitude" db:"longitude"`
	Active    bool      `json:"active" db:"active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type EventRepository interface {
	Create(*Event) error
	FindAll() ([]*Event, error)
	FindByID(int) (*Event, error)
	FindByOwnerID(int) ([]*Event, error)
	FindByName(string) ([]*Event, error)
	FindByLocation(string) ([]*Event, error)
	Update(*Event) error
	DeActivate(int) error
	ReActivate(int) error
	Checkin(int, int) error
}
