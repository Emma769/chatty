package psql

import (
	"context"

	"github.com/emma769/chatty/internal/model"
	"github.com/emma769/chatty/internal/repository"
)

type CreateUserParam struct {
	Username string
	Email    string
	Password []byte
}

func (q *Queries) CreateUser(ctx context.Context, param CreateUserParam) (*model.User, error) {
	query := `
  INSERT INTO users (username, email, password) VALUES ($1, $2, $3)
  RETURNING user_id, username, email, password,
  version, created_at, updated_at, deleted_at;
  `

	row := q.db.QueryRowContext(ctx, query, param.Username, param.Email, param.Password)

	user := &model.User{}

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
