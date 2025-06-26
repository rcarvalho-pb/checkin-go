package participant

import "context"

type participantStorage interface {
	Create(context.Context, *Participant) (int, error)
	List(context.Context) ([]*Participant, error)
	FindById(context.Context, int) (*Participant, error)
	FindByEmail(context.Context, string) (*Participant, error)
	DeleteById(context.Context, int) error
}
