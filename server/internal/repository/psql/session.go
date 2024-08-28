package psql

import (
	"context"
	"time"

	"github.com/emma769/chatty/internal/data"
)

type CreateSessionParam struct {
	Scope     data.Scope
	UserID    string
	Email     string
	Hash      []byte
	IsRevoked bool
	ValidTill time.Time
}

func (q *Queries) CreateSession(ctx context.Context, param CreateSessionParam) error {
	query := `
  INSERT INTO sessions (scope, hash, valid_till, is_revoked, user_id, email) VALUES ($1, $2, $3, $4, $5, $6);
  `

	_, err := q.db.ExecContext(
		ctx,
		query,
		param.Scope,
		param.Hash,
		param.ValidTill,
		param.IsRevoked,
		param.UserID,
		param.Email,
	)
	if err != nil {
		return err
	}

	return nil
}
