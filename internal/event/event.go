package event

import "time"

type Event struct {
	ID        int
	Name      string
	Location  string
	StartsAt  time.Time
	EndsAt    time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EventRepository interface {
	Create(*Event) error
	FindAll() ([]*Event, error)
	FindByID(int) (*Event, error)
	FindByName(string) ([]*Event, error)
	FindByLocation(string) ([]*Event, error)
	Update(*Event) error
	DeActivate(int) error
	ReActivate(int) error
}
