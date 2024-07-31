package passlib

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Passlib struct {
	cost int
}

func New() *Passlib {
	return &Passlib{bcrypt.DefaultCost}
}

func (p Passlib) Hash(pl string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pl), p.cost)
}

func (p Passlib) Verify(pl string, hash []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, []byte(pl))

	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
