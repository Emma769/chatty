package psql

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
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

type getter interface {
	Get(string, string) string
	GetInt(string, int) int
}

func NewRepository(cfg getter) (*Repository, error) {
	uri := cfg.Get("POSTGRES_URI", "")

	if uri == "" {
		return nil, errors.New("postgres uri cannot be blank")
	}

	url, err := pq.ParseURL(uri)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(DRIVER_NAME, url)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.GetInt("DB_MAX_OPEN_CONNS", 25))
	db.SetMaxIdleConns(cfg.GetInt("DB_MAX_IDLE_CONNS", 25))

	maxIdleTime := cfg.GetInt("DB_CONN_MAX_IDLE_TIME", 15)
	db.SetConnMaxIdleTime(time.Duration(maxIdleTime) * time.Second)

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
