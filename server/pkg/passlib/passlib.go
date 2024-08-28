package passlib

import "golang.org/x/crypto/bcrypt"

func Hash(plain string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
}

func IsMatch(plain string, hash []byte) bool {
	return bcrypt.CompareHashAndPassword(hash, []byte(plain)) == nil
}
