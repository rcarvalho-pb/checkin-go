package storage

import (
	"context"

	"github.com/rcarvalho-pb/checkin-go/internal/participant"
)

type ParticipantStorage interface {
	Create(context.Context, *participant.Participant) (int, error)
	List(context.Context) ([]*participant.Participant, error)
	FindById(context.Context, int) (*participant.Participant, error)
	FindByEmail(context.Context, string) (*participant.Participant, error)
	DeleteById(context.Context, int) error
}
