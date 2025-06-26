package postgres

import (
	"context"
	"fmt"

	"github.com/rcarvalho-pb/checkin-go/internal/model"
)

type eventStorage struct {
	*DB
}

func NewEventStorage(db *DB) *eventStorage {
	return &eventStorage{
		db,
	}
}

func (es *eventStorage) Create(ctx context.Context, e *model.Event) (int, error) {
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

func (es *eventStorage) List(ctx context.Context) ([]*model.Event, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()
	query := `
	SELECT * FROM
		events;
	`
	var events []*model.Event
	if err := es.SelectContext(ctx, &events, query); err != nil {
		return nil, fmt.Errorf("error selecting from events table: %w", err)
	}
	return events, nil
}

func (es *eventStorage) FindById(ctx context.Context, id int) (*model.Event, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()
	query := `
	SELECT * FROM
		events
	WHERE 
		id = :id;
	`
	var event *model.Event
	if err := es.GetContext(ctx, &event, query, id); err != nil {
		return nil, err
	}
	return event, nil
}

func (es *eventStorage) DeleteById(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()
	stmt := `
	DELETE from
		events
	WHERE
	 id = :id;
	`
	if _, err := es.ExecContext(ctx, stmt, id); err != nil {
		return err
	}
	return nil
}
