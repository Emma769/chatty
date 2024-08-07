package psql

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"

	"github.com/emma769/chatty/internal/config"
)

const DRIVER_NAME = "postgres"

type dbtx interface {
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...any) *sql.Row
	ExecContext(context.Context, string, ...any) (sql.Result, error)
}

type Queries struct {
	db dbtx
}

func newQueries(db dbtx) *Queries {
	return &Queries{db}
}

type Repository struct {
	*Queries
	db *sql.DB
}

func NewRepository(cfg *config.Config) (*Repository, error) {
	if cfg.POSTGRES_URI == "" {
		return nil, errors.New("postgres uri cannot be blank")
	}

	url, err := pq.ParseURL(cfg.POSTGRES_URI)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(DRIVER_NAME, url)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.DB_MAX_OPEN_CONNS)
	db.SetMaxIdleConns(cfg.DB_MAX_IDLE_CONNS)
	db.SetConnMaxIdleTime(time.Duration(cfg.DB_CONN_MAX_IDLE_TIME) * time.Second)

	repository := &Repository{
		newQueries(db),
		db,
	}

	return repository, nil
}

func (r *Repository) Ping() error {
	return r.db.Ping()
}

func (r *Repository) Close() error {
	return r.db.Close()
}
