package postgres

import (
	"context"
	"fmt"

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
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()
	stmt := `
	INSERT INTO events
		(name, location, starts_at, ends_at)
	VALUES
		(:name, :location, :starts_at, :ends_at)
	RETURNING id;
	`
	row, err := ps.NamedQueryContext(ctx, stmt, p)
	if err != nil {
		return 0, fmt.Errorf("named insert error: %w", err)
	}
	defer row.Close()
	var id int
	if row.Next() {
		if err := row.Scan(&id); err != nil {
			return 0, fmt.Errorf("scan error: %w", err)
		}
	}
	return id, nil
}

func (ps *participantStorage) List(ctx context.Context) ([]*participant.Participant, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()
	query := `
	SELECT * FROM
		events;
	`
	var p []*participant.Participant
	if err := ps.SelectContext(ctx, &p, query); err != nil {
		return nil, fmt.Errorf("error selecting from events table: %w", err)
	}
	return p, nil
}

func (ps *participantStorage) FindById(ctx context.Context, id int) (*participant.Participant, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()
	query := `
	SELECT * FROM
		events
	WHERE 
		id = :id;
	`
	var p *participant.Participant
	if err := ps.GetContext(ctx, &p, query, id); err != nil {
		return nil, err
	}
	return p, nil
}

func (ps *participantStorage) FindByEmail(ctx context.Context, email string) (*participant.Participant, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()
	query := `
	SELECT * FROM
		events
	WHERE 
		email = :email;
	`
	var p *participant.Participant
	if err := ps.GetContext(ctx, &p, query, email); err != nil {
		return nil, err
	}
	return p, nil
}

func (ps *participantStorage) DeleteById(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()
	stmt := `
	DELETE from
		events
	WHERE
	 id = :id;
	`
	if _, err := ps.ExecContext(ctx, stmt, id); err != nil {
		return err
	}
	return nil
}
