package session

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"time"

	"github.com/emma769/chatty/internal/data"
	"github.com/emma769/chatty/internal/repository/psql"
)

var timeout = 3 * time.Second

type storer interface {
	CreateSession(context.Context, psql.CreateSessionParam) error
}

type Service struct {
	timeout time.Duration
	store   storer
}

func NewService(store storer) *Service {
	return &Service{
		timeout: timeout,
		store:   store,
	}
}

func (s *Service) Create(ctx context.Context, in data.SessionIn) (string, time.Time, error) {
	b := make([]byte, 32)

	if _, err := rand.Read(b); err != nil {
		return "", time.Time{}, err
	}

	token := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(b)

	sum := sha256.Sum256([]byte(token))
	hash := sum[:]

	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	validtill := time.Now().Add(in.ValidFor)

	err := s.store.CreateSession(ctx, psql.CreateSessionParam{
		UserID:    in.UserID,
		Email:     in.Email,
		Hash:      hash,
		ValidTill: validtill,
		Scope:     in.Scope,
	})
	if err != nil {
		return "", time.Time{}, err
	}

	return token, validtill, nil
}
