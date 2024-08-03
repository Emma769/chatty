package funclib

import (
	"cmp"
	"fmt"
	"math/rand"
	"net/mail"
	"strings"
)

func ValidEmail(s string) bool {
	if s == "" {
		return false
	}

	_, err := mail.ParseAddress(s)
	return err == nil
}

func Gte[T cmp.Ordered](a, b T) bool { return a >= b }
func Lte[T cmp.Ordered](a, b T) bool { return a <= b }

func AsciiLower() string {
	var b strings.Builder

	for i := 'a'; i < 'a'+26; i++ {
		b.WriteRune(i)
	}

	return b.String()
}

func RandString(n int) string {
	pl := AsciiLower()
	var b strings.Builder

	for range n {
		b.WriteByte(pl[rand.Intn(len(pl)-1)])
	}

	return strings.Join(Shuffle(strings.Split(b.String(), "")), "")
}

func Shuffle[T any](ts []T) []T {
	for i := range ts {
		j := rand.Intn(i + 1)
		ts[i], ts[j] = ts[j], ts[i]
	}

	return ts
}

func RandInt(mn, mx int) int {
	return rand.Intn(mx-mn) + mn
}

func RandName() string {
	s := RandString(RandInt(5, 10))
	return strings.ToUpper(string(s[0])) + string(s[1:])
}

var domains = []string{"gmail", "yahoo", "hotmail", "msn"}

func RandEmail() string {
	return fmt.Sprintf("%s@%s.com", RandString(7), domains[rand.Intn(len(domains)-1)])
}
