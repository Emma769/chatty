package handler

import (
	"context"

	"github.com/emma769/chatty/internal/model"
)

type userSvc interface {
	Create(context.Context, model.UserIn) (*model.User, error)
}

type Service struct {
	User userSvc
}

type Handler struct {
	*Service
}

func New(svc *Service) *Handler {
	return &Handler{svc}
}
