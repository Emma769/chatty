package user

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/emma769/chatty/internal/model"
	"github.com/emma769/chatty/internal/repository"
	"github.com/emma769/chatty/internal/repository/psql"
	"github.com/emma769/chatty/pkg/funclib"
)

type storeStub struct {
	err error
}

func (s storeStub) CreateUser(
	ctx context.Context,
	param psql.CreateUserParam,
) (*model.User, error) {
	user := &model.User{
		Username: param.Username,
		Email:    param.Email,
		Password: param.Password,
	}

	return user, s.err
}

type hashStub struct {
	err error
}

func (h hashStub) Hash(plain string) ([]byte, error) {
	return []byte(plain), h.err
}

func TestService_Create(t *testing.T) {
	type arg struct {
		name string
		ctx  context.Context
		svc  *Service
		in   model.UserIn
	}

	for range 5 {
		tc := arg{
			name: fmt.Sprintf("%s_create", funclib.RandString(10)),
			ctx:  context.Background(),
			svc:  NewService(&storeStub{}, &hashStub{}),
			in: model.UserIn{
				Username: funclib.RandName(),
				Email:    funclib.RandEmail(),
				Password: funclib.RandString(8),
			},
		}

		t.Run(tc.name, func(t *testing.T) {
			user, err := tc.svc.Create(tc.ctx, tc.in)
			require.NoError(t, err)
			require.NotEmpty(t, user)
			require.Equal(t, user.Username, tc.in.Username)
			require.Equal(t, user.Email, tc.in.Email)
		})
	}
}

func TestService_Create_ReturnsHasherError(t *testing.T) {
	type arg struct {
		name string
		ctx  context.Context
		svc  *Service
		in   model.UserIn
	}

	for range 3 {
		tc := arg{
			name: fmt.Sprintf("%s_create", funclib.RandString(10)),
			ctx:  context.Background(),
			svc:  NewService(&storeStub{}, &hashStub{errors.New("invalid arg")}),
			in: model.UserIn{
				Username: funclib.RandName(),
				Email:    funclib.RandEmail(),
				Password: funclib.RandString(8),
			},
		}

		t.Run(tc.name, func(t *testing.T) {
			user, err := tc.svc.Create(tc.ctx, tc.in)
			require.Error(t, err)
			require.Nil(t, user)
		})
	}
}

func TestService_Create_ReturnsStorerError(t *testing.T) {
	type arg struct {
		name string
		ctx  context.Context
		svc  *Service
		in   model.UserIn
	}

	for range 3 {
		tc := arg{
			name: fmt.Sprintf("%s_create", funclib.RandString(10)),
			ctx:  context.Background(),
			svc:  NewService(&storeStub{repository.ErrDuplicateKey}, &hashStub{}),
			in: model.UserIn{
				Username: funclib.RandName(),
				Email:    funclib.RandEmail(),
				Password: funclib.RandString(8),
			},
		}

		t.Run(tc.name, func(t *testing.T) {
			user, err := tc.svc.Create(tc.ctx, tc.in)
			require.Error(t, err)
			require.Nil(t, user)
		})
	}
}
