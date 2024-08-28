package data

import "time"

type Scope string

const (
	Scope_Authentication Scope = "authentication"
)

type SessionIn struct {
	UserID   string
	Email    string
	Scope    Scope
	ValidFor time.Duration
}
