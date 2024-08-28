package tokens

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto/v2"
)

type Maker struct {
	key    []byte
	paseto *paseto.V2
}

func NewMaker(key string) (*Maker, error) {
	if len(key) != 32 {
		return nil, fmt.Errorf("invalid symmetric key - must contain 32 characters")
	}

	return &Maker{
		key:    []byte(key),
		paseto: paseto.NewV2(),
	}, nil
}

func (m *Maker) Encrypt(userId string, validtill time.Duration) (string, time.Time, error) {
	now := time.Now()
	validfor := now.Add(validtill)

	jwttoken := paseto.JSONToken{
		IssuedAt:   now,
		NotBefore:  now,
		Expiration: validfor,
	}

	jwttoken.Set("payload", NewPayload(userId))

	token, err := m.paseto.Encrypt(m.key, jwttoken, nil)
	if err != nil {
		return "", time.Time{}, err
	}

	return token, validfor, nil
}
