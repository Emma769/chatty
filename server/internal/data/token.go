package data

import "time"

type AccessToken struct {
	Value     string    `json:"value"`
	ValidTill time.Time `json:"valid_till"`
}

type RefreshToken struct {
	Value     string    `json:"value"`
	ValidTill time.Time `json:"valid_till"`
}

type Token struct {
	User         *User         `json:"user"`
	AccessToken  AccessToken   `json:"access_token"`
	RefreshToken *RefreshToken `json:"refresh_token,omitempty"`
}

type TokenIn struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
