package repository

import (
	"context"
	"errors"
	"time"

	"github.com/IamStubborN/otus-golang/2019_08_27_homework_16/calendar/config"
	"github.com/IamStubborN/otus-golang/2019_08_27_homework_16/calendar/service/event/domain"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	pool *sqlx.DB
}

func NewDatabase(cfg *config.Config) (EvInterface, error) {
	conn, err := initialSQLConn(cfg.Storage)
	if err != nil {
		return nil, err
	}

	return &Database{
		pool: conn,
	}, nil
}

func (d *Database) Close() error {
	if err := d.pool.Close(); err != nil {
		return err
	}

	return nil
}

func initialSQLConn(cfg config.Storage) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.DSN)
	if err != nil {
		return nil, err
	}

	if err := retryConnect(db, cfg.Retry); err != nil {
		return nil, err
	}

	return db, nil
}

func retryConnect(db *sqlx.DB, fatalRetry int) error {
	var retryCount int
	for range time.NewTicker(time.Second).C {
		retryCount++
		if err := db.Ping(); err == nil {
			break
		}

		if fatalRetry == retryCount {
			return errors.New("can't connect to database")
		}
	}

	return nil
}

func (d *Database) Create(ctx context.Context, ev *domain.Event) (*domain.Event, error) {
	query := `
	insert into events("name", description, "date") 
		values (:name, :description, :date) returning id
	`

	argsQ := map[string]interface{}{
		"name":        ev.Name,
		"description": ev.Description,
		"date":        ev.Date,
	}

	stmt, err := d.pool.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var id uint64
	err = stmt.QueryRowxContext(ctx, argsQ).Scan(&id)
	if err != nil {
		return nil, err
	}

	ev.ID = id

	return ev, nil
}

func (d *Database) Read(ctx context.Context, eventID uint64) (*domain.Event, error) {
	query := `select id, "name", description, "date" from events where id=$1`

	var ev domain.Event

	err := d.pool.QueryRowxContext(ctx, query, eventID).Scan(&ev.ID, &ev.Name, &ev.Description, &ev.Date)
	if err != nil {
		return nil, err
	}

	return &ev, nil
}

func (d *Database) Update(ctx context.Context, ev *domain.Event) (bool, error) {
	query := `update events set "name"=:name, description=:description, "date"=:date 
		where id=:id`

	argsQ := map[string]interface{}{
		"id":          ev.ID,
		"name":        ev.Name,
		"description": ev.Description,
		"date":        ev.Date,
	}

	res, err := d.pool.NamedExecContext(ctx, query, argsQ)
	if err != nil {
		return false, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, errors.New("can't update event")
	}

	return true, nil
}

func (d *Database) Delete(ctx context.Context, eventID uint64) (bool, error) {
	query := `delete from events where id=$1`

	res, err := d.pool.ExecContext(ctx, query, eventID)
	if err != nil {
		return false, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, errors.New("can't delete event")
	}

	return true, nil
}
