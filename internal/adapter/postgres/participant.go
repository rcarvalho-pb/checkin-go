package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/rcarvalho-pb/checkin-go/internal/participant"
)

type PostgresParticipant struct {
	*sqlx.DB
}

func OpenDB(dsn string) participant.ParticipantRepository {
	conn := connectToDB(dsn)
	if conn == nil {
		log.Panic("can't connect to database")
	}
	return &PostgresParticipant{conn}
}

func (pp *PostgresParticipant) Create(p *participant.Participant) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `
	INSERT INTO
		participants
		(name, email, password, role)
	VALUES
		(:name, :email, :password, :role)
	`
	if _, err := pp.NamedExecContext(ctx, stmt, p); err != nil {
		return err
	}
	return nil
}

func (pp *PostgresParticipant) FindByID(id int) (*participant.Participant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
	SELECT * FROM
		participants
	WHERE
		id = :id;
	`
	var p *participant.Participant
	if err := pp.GetContext(ctx, p, query, id); err != nil {
		return nil, err
	}
	return p, nil
}

func (pp *PostgresParticipant) FindByEmail(email string) (*participant.Participant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
	SELECT * FROM
		participants
	WHERE
		email = $1
	`
	var p participant.Participant
	if err := pp.GetContext(ctx, &p, query, email); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &p, nil
}

func (pp *PostgresParticipant) Update(p *participant.Participant) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `
	UPDATE
		participants
	SET
	name = :name, email = :email, password = :password, active = :active, role = :role, created_at = :created_at, updated_at = :updated_at
	WHERE
		id = :id;
	`
	if _, err := pp.NamedExecContext(ctx, stmt, p); err != nil {
		return err
	}
	return nil
}

func (pp *PostgresParticipant) ReActivate(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `
	UPDATE
		participants
	SET
		active = true 
	WHERE
		id = :id;
	`
	if _, err := pp.NamedExecContext(ctx, stmt, id); err != nil {
		return err
	}
	return nil
}

func (pp *PostgresParticipant) DeActivate(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `
	UPDATE
		participants
	SET
		active = false
	WHERE
		id = :id;
	`
	if _, err := pp.NamedExecContext(ctx, stmt, id); err != nil {
		return err
	}
	return nil
}
