package handler

import (
	"context"
	"log/slog"
	"time"

	"github.com/emma769/chatty/internal/data"
)

type userService interface {
	Create(context.Context, data.UserIn) (*data.User, error)
	FindByEmail(context.Context, string) (*data.User, error)
}

type sessionService interface {
	Create(context.Context, data.SessionIn) (string, time.Time, error)
}

type Service struct {
	User    userService
	Session sessionService
}

type tokenMaker interface {
	Encrypt(string, time.Duration) (string, time.Time, error)
}

type getter interface {
	GetDuration(string) time.Duration
}

type Handler struct {
	*Service
	ctx    context.Context
	Getter getter
	Logger *slog.Logger
	Token  tokenMaker
}

func NewWithService(
	ctx context.Context,
	getter getter,
	token tokenMaker,
	logger *slog.Logger,
	service *Service,
) *Handler {
	return &Handler{
		Service: service,
		ctx:     ctx,
		Getter:  getter,
		Logger:  logger,
		Token:   token,
	}
}
