package psql

import (
	"cmp"
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

var driver = "postgres"

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
	return &Queries{db: db}
}

type Repository struct {
	*Queries
	db *sql.DB
}

type getter interface {
	GetInt(string) int
	GetDuration(string) time.Duration
}

func NewRepository(uri string, getter getter) (*Repository, error) {
	url, err := pq.ParseURL(uri)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(driver, url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cmp.Or(getter.GetInt("DB_MAX_OPEN_CONNS"), 25))
	db.SetMaxIdleConns(cmp.Or(getter.GetInt("DB_MAX_IDLE_CONNS"), 25))
	db.SetConnMaxIdleTime(cmp.Or(getter.GetDuration("DB_CONN_MAX_IDLE_TIME"), 15*time.Second))

	return &Repository{
		Queries: newQueries(db),
		db:      db,
	}, nil
}

func (r *Repository) Close() error {
	return r.db.Close()
}
