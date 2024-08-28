package repository

import (
	"errors"
	"strings"
)

var (
	ErrDuplicateKey = errors.New("duplicate key")
	ErrNotFound     = errors.New("not found")
)

func DuplicateKey(err error) bool {
	if err == nil {
		return false
	}

	return strings.Contains(err.Error(), "duplicate")
}
