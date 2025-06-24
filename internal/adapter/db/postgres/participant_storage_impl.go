package postgres

import (
	"context"

	"github.com/rcarvalho-pb/checkin-go/internal/participant"
)

type participantStorage struct {
	*DB
}

// Create(context.Context, *participant.Participant) (int, error)
// List(context.Context) ([]*participant.Participant, error)
// FindById(context.Context, int) (*participant.Participant, error)
// FindByEmail(context.Context, string) (*participant.Participant, error)
// DeleteById(context.Context, int) error

func NewParticipantStorage(db *DB) *participantStorage {
	return &participantStorage{
		db,
	}
}

func (ps *participantStorage) Create(ctx context.Context, p *participant.Participant) (int, error) {
	return 0, nil
}
