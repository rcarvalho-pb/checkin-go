package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rcarvalho-pb/checkin-go/internal/event"
)

type PostgresEvent struct {
	DB *sqlx.DB
}

func (pe *PostgresEvent) Create(e *event.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		INSERT INTO events (name, location, starts_at, ends_at, owner_id, latitude, longitude, active, created_at, updated_at)
	VALUES (:name, :location, :starts_at, :ends_at, :owner_id, :latitude, :longitude, :active, :created_at, :updated_at)
		RETURNING id;
	`
	stmt, err := pe.DB.PrepareNamedContext(ctx, query)
	if err != nil {
		return err
	}
	return stmt.GetContext(ctx, &e.ID, e)
}

func (pe *PostgresEvent) FindAll() ([]*event.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var events []*event.Event
	query := `SELECT * FROM events ORDER BY starts_at`
	if err := pe.DB.SelectContext(ctx, &events, query); err != nil {
		return nil, err
	}
	return events, nil
}

func (pe *PostgresEvent) FindByID(id int) (*event.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var e event.Event
	query := `SELECT * FROM events WHERE id = $1`
	if err := pe.DB.GetContext(ctx, &e, query, id); err != nil {
		return nil, err
	}
	return &e, nil
}

func (pe *PostgresEvent) FindByName(name string) ([]*event.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var events []*event.Event
	query := `SELECT * FROM events WHERE name ILIKE '%' || $1 || '%'`
	if err := pe.DB.SelectContext(ctx, &events, query, name); err != nil {
		return nil, err
	}
	return events, nil
}

func (pe *PostgresEvent) FindByOwnerID(ownerID int) ([]*event.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var events []*event.Event
	query := `SELECT * FROM events WHERE owner_id = $1 ORDER BY starts_at`
	if err := pe.DB.SelectContext(ctx, &events, query, ownerID); err != nil {
		return nil, err
	}
	return events, nil
}

func (pe *PostgresEvent) FindByLocation(location string) ([]*event.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var events []*event.Event
	query := `SELECT * FROM events WHERE location ILIKE '%' || $1 || '%'`
	if err := pe.DB.SelectContext(ctx, &events, query, location); err != nil {
		return nil, err
	}
	return events, nil
}

func (pe *PostgresEvent) Update(e *event.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		UPDATE events
		SET name = :name,
			location = :location,
			starts_at = :starts_at,
			ends_at = :ends_at,
			active = :active,
			updated_at = :updated_at
		WHERE id = :id
	`
	_, err := pe.DB.NamedExecContext(ctx, query, e)
	return err
}

func (pe *PostgresEvent) DeActivate(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `UPDATE events SET active = false, updated_at = now() WHERE id = $1`
	_, err := pe.DB.ExecContext(ctx, query, id)
	return err
}

func (pe *PostgresEvent) ReActivate(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `UPDATE events SET active = true, updated_at = now() WHERE id = $1`
	_, err := pe.DB.ExecContext(ctx, query, id)
	return err
}

func (pe *PostgresEvent) Checkin(eventID, participantID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
	INSERT INTO event_participants (event_id, participant_id, checked_in)
	VALUES ($1, $2, $3)
	`
	stmt, err := pe.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, eventID, participantID, true)
	return err
}
