package participant

import (
	"time"

	participant_role "github.com/rcarvalho-pb/checkin-go/internal/participant/roles"
)

type Participant struct {
	ID        int                   `json:"id"`
	Name      string                `json:"name"`
	Email     string                `json:"email"`
	Password  string                `json:"password,omitempty"`
	Active    bool                  `json:"active,omitempty"`
	Role      participant_role.Role `json:"role"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

type ParticipantRepository interface {
	Create(*Participant) error
	FindByID(int) (*Participant, error)
	FindByEmail(string) (*Participant, error)
	Update(*Participant) error
	DeActivate(int) error
	ReActivate(int) error
}
