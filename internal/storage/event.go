package storage

import (
	"context"

	"github.com/rcarvalho-pb/checkin-go/internal/model"
)

type EventStorage interface {
	Create(context.Context, *model.Event) (int, error)
	List(context.Context) ([]*model.Event, error)
	FindById(context.Context, int) (*model.Event, error)
	DeleteById(context.Context, int) error
}
