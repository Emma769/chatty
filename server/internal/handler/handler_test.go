package handler

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/emma769/chatty/internal/config"
	"github.com/emma769/chatty/internal/repository/psql"
	"github.com/emma769/chatty/internal/services/user"
	"github.com/emma769/chatty/pkg/passlib"
)

func newTestHandler(t *testing.T) *Handler {
	repo, err := psql.NewRepository(&config.Getter{})
	if err != nil {
		require.FailNow(t, err.Error())
	}

	userService := user.NewService(repo, passlib.New())

	handlerService := &Service{
		User: userService,
	}

	return New(handlerService)
}
