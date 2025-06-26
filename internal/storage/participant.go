package storage

import (
	"context"

	"github.com/rcarvalho-pb/checkin-go/internal/model"
)

type ParticipantStorage interface {
	Create(context.Context, *model.Participant) (int, error)
	List(context.Context) ([]*model.Participant, error)
	FindById(context.Context, int) (*model.Participant, error)
	FindByEmail(context.Context, string) (*model.Participant, error)
	DeleteById(context.Context, int) error
}
