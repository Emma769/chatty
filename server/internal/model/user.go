package model

import (
	"time"

	"github.com/emma769/chatty/pkg/funclib"
	"github.com/emma769/chatty/pkg/validator"
)

type User struct {
	UserID    string     `json:"user_id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  []byte     `json:"-"`
	Version   int64      `json:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"-"`
}

type UserIn struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (in UserIn) Validate() error {
	return validator.Check(
		validator.New(),
		in,
		func(i UserIn) (bool, string) { return i.Username != "", "username:cannot be blank" },
		func(i UserIn) (bool, string) { return i.Email != "", "email:cannot be blank" },
		func(i UserIn) (bool, string) { return i.Password != "", "password:cannot be blank" },
		func(i UserIn) (bool, string) {
			return funclib.ValidEmail(i.Email), "email:provide a valid email"
		},
		func(i UserIn) (bool, string) {
			return funclib.Gte(len(i.Password), 8), "password:cannot be less than 8 characters"
		},
		func(i UserIn) (bool, string) {
			return funclib.Lte(len(i.Password), 50), "password:cannot be more 50 characters"
		},
	)
}
