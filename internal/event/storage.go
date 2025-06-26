package event

import "context"

type eventStorage interface {
	Create(context.Context, *Event) (int, error)
	List(context.Context) ([]*Event, error)
	FindById(context.Context, int) (*Event, error)
	DeleteById(context.Context, int) error
}
