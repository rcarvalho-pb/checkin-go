package postgres

import (
	"context"
	"fmt"

	"github.com/rcarvalho-pb/checkin-go/internal/event"
)

type eventStorage struct {
	*DB
}

func NewEventStore(db *DB) *eventStorage {
	return &eventStorage{
		db,
	}
}

// Create(context.Context, *event.Event) (int, error)
// List(context.Context) ([]*event.Event, error)
// FindById(context.Context, int) (*event.Event, error)
// DeleteById(context.Context, int) error

func (es *eventStorage) Create(ctx context.Context, e *event.Event) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()
	stmt := `
	INSERT INTO events
		(name, location, starts_at, ends_at)
	VALUES
		(:name, :location, :starts_at, :ends_at)
	RETURNING id;
	`
	row, err := es.NamedQueryContext(ctx, stmt, e)
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

func (es *eventStorage) List(ctx context.Context) ([]*event.Event, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()
	query := `
	SELECT * FROM
		participants;
	`
	var events []*event.Event
	if err := es.SelectContext(ctx, &events, query); err != nil {
		return nil, fmt.Errorf("error selecting from events table: %w", err)
	}
	return events, nil
}

func (es *eventStorage) FindById(ctx context.Context, id int) (*event.Event, error) {

}
