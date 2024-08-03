package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"

	"github.com/emma769/chatty/internal/model"
	"github.com/emma769/chatty/pkg/funclib"
)

func TestHandler_CreateUser(t *testing.T) {
	router := chi.NewRouter()
	handler := newTestHandler(t)
	handler.Register(router)

	server := httptest.NewServer(router)

	client := UserClient{
		url: server.URL,
	}

	in := model.UserIn{
		Username: funclib.RandName(),
		Email:    funclib.RandEmail(),
		Password: funclib.RandString(8),
	}

	user, err := client.CreateUser(in)
	require.NoError(t, err)
	require.NotEmpty(t, user)
}
