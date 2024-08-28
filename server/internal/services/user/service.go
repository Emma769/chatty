package user

import (
	"context"
	"errors"
	"time"

	"github.com/emma769/chatty/internal/data"
	"github.com/emma769/chatty/internal/repository"
	"github.com/emma769/chatty/internal/repository/psql"
	"github.com/emma769/chatty/internal/services"
	"github.com/emma769/chatty/pkg/passlib"
)

var timeout = 3 * time.Second

type store interface {
	CreateUser(context.Context, psql.CreateUserParam) (*data.User, error)
	FindUserByEmail(context.Context, string) (*data.User, error)
}

type Service struct {
	timeout time.Duration
	store   store
}

func NewService(store store) *Service {
	return &Service{
		store:   store,
		timeout: timeout,
	}
}

func (s *Service) Create(ctx context.Context, in data.UserIn) (*data.User, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	password, err := passlib.Hash(in.Password)
	if err != nil {
		return nil, err
	}

	user, err := s.store.CreateUser(ctx, psql.CreateUserParam{
		Username: in.Username,
		Email:    in.Email,
		Password: password,
	})
	if err != nil && errors.Is(err, repository.ErrDuplicateKey) {
		return nil, services.ErrDuplicateKey
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) FindByEmail(ctx context.Context, email string) (*data.User, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	user, err := s.store.FindUserByEmail(ctx, email)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, services.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}
