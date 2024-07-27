package funclib

import (
	"cmp"
	"net/mail"
)

func ValidEmail(s string) bool {
	if s == "" {
		return false
	}

	_, err := mail.ParseAddress(s)
	return err == nil
}

func Gte[T cmp.Ordered](a, b T) bool { return a > b }
func Lte[T cmp.Ordered](a, b T) bool { return a < b }
