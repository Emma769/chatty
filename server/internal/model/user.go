package model

import "time"

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
	Username string `json:"username" validate:"required"`
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
