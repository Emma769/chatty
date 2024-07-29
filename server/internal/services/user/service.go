package user

import (
	"context"
	"errors"
	"time"

	"github.com/emma769/chatty/internal/model"
	"github.com/emma769/chatty/internal/repository"
	"github.com/emma769/chatty/internal/repository/psql"
	"github.com/emma769/chatty/internal/services"
)

type store interface {
	CreateUser(context.Context, psql.CreateUserParam) (*model.User, error)
}

type Service struct {
	timeout time.Duration
	store   store
}

func NewService(store store) *Service {
	return &Service{3 * time.Second, store}
}

func (s *Service) Create(ctx context.Context, in model.UserIn) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	param := psql.CreateUserParam{
		Username: in.Username,
		Email:    in.Email,
		Password: []byte{},
	}

	user, err := s.store.CreateUser(ctx, param)

	if err != nil && errors.Is(err, repository.ErrDuplicateKey) {
		return nil, services.ErrDuplicateKey
	}

	return user, err
}
