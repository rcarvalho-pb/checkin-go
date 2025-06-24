package storage

import (
	"context"

	"github.com/rcarvalho-pb/checkin-go/internal/event"
)

type EvetStorage interface {
	Create(context.Context, *event.Event) (int, error)
	List(context.Context) ([]*event.Event, error)
	FindById(context.Context, int) (*event.Event, error)
	DeleteById(context.Context, int) error
}
