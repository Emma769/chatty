package psql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/emma769/chatty/internal/data"
	"github.com/emma769/chatty/internal/repository"
)

type CreateUserParam struct {
	Username string
	Email    string
	Password []byte
}

func (q *Queries) CreateUser(ctx context.Context, param CreateUserParam) (*data.User, error) {
	query := `
  INSERT INTO users (username, email, password) VALUES ($1, $2, $3)
  RETURNING user_id, username, email, password, version, created_at, updated_at, deleted_at;
  `

	row := q.db.QueryRowContext(ctx, query, param.Username, param.Email, param.Password)

	user := new(data.User)

	err := row.Scan(
		&user.UserID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Version,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	if err != nil && repository.DuplicateKey(err) {
		return nil, repository.ErrDuplicateKey
	}

	return user, err
}

func (q *Queries) FindUserByEmail(ctx context.Context, email string) (*data.User, error) {
	query := `
  SELECT user_id, username, email, password, version, created_at, updated_at, deleted_at
  FROM users WHERE email = $1;
  `

	row := q.db.QueryRowContext(ctx, query, email)

	user := new(data.User)

	err := row.Scan(
		&user.UserID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Version,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, repository.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}
