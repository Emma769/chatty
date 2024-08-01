package psql

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/emma769/chatty/internal/model"
	"github.com/emma769/chatty/internal/repository"
	"github.com/emma769/chatty/pkg/funclib"
)

func TestQueries_CreateUser(t *testing.T) {
	type arg struct {
		name  string
		ctx   context.Context
		repo  *Repository
		param CreateUserParam
	}

	for range 3 {
		tc := arg{
			name: fmt.Sprintf("%s_create_user", funclib.RandString(10)),
			ctx:  context.Background(),
			repo: newTestRepo(t),
			param: CreateUserParam{
				Username: funclib.RandName(),
				Email:    funclib.RandEmail(),
				Password: []byte(funclib.RandString(10)),
			},
		}

		t.Run(tc.name, func(t *testing.T) {
			user, err := tc.repo.CreateUser(tc.ctx, tc.param)
			require.NoError(t, err)
			require.NotEmpty(t, user)
			require.Equal(t, user.Username, tc.param.Username)
			require.Equal(t, user.Email, tc.param.Email)
			require.NotEmpty(t, user.CreatedAt)
		})
	}
}

func createUser(t *testing.T, n int) []*model.User {
	users := []*model.User{}

	type arg struct {
		name  string
		ctx   context.Context
		repo  *Repository
		param CreateUserParam
	}

	for range n {
		tc := arg{
			name: fmt.Sprintf("%s_create_user", funclib.RandString(10)),
			ctx:  context.Background(),
			repo: newTestRepo(t),
			param: CreateUserParam{
				Username: funclib.RandName(),
				Email:    funclib.RandEmail(),
				Password: []byte(funclib.RandString(10)),
			},
		}

		user, err := tc.repo.CreateUser(tc.ctx, tc.param)
		if err != nil {
			t.FailNow()
		}

		users = append(users, user)
	}

	return users
}

func TestQueries_CreateUser_ReturnsDuplicateKeyError(t *testing.T) {
	users := createUser(t, 3)

	type arg struct {
		name  string
		ctx   context.Context
		repo  *Repository
		param CreateUserParam
	}

	for _, user := range users {
		tc := arg{
			name: fmt.Sprintf("%s_create_user", funclib.RandString(10)),
			ctx:  context.Background(),
			repo: newTestRepo(t),
			param: CreateUserParam{
				Username: user.Username,
				Email:    user.Email,
				Password: user.Password,
			},
		}

		t.Run(tc.name, func(t *testing.T) {
			u, err := tc.repo.CreateUser(tc.ctx, tc.param)
			require.Error(t, err)
			require.Nil(t, u)
			require.ErrorIs(t, err, repository.ErrDuplicateKey)
		})
	}
}
