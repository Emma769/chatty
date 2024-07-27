package validator

import (
	"fmt"
	"strings"
	"sync"
)

type valerr struct {
	*sync.RWMutex
	err map[string]string
}

func (e valerr) Error() string {
	var b strings.Builder

	for k, v := range e.err {
		b.WriteString(fmt.Sprintln(k, v))
	}

	return b.String()
}

type validator struct {
	*valerr
}

func New() *validator {
	return &validator{&valerr{
		new(sync.RWMutex),
		map[string]string{},
	}}
}

func (v *validator) add(key, msg string) {
	v.Lock()
	defer v.Unlock()

	_, ok := v.err[key]

	if !ok {
		v.err[key] = msg
	}
}

func (v validator) valid() bool {
	return len(v.err) == 0
}

type validationfn[T any] func(T) (bool, string)

func Check[T any](v *validator, t T, fns ...validationfn[T]) error {
	for _, fn := range fns {
		ok, errmsg := fn(t)

		if !ok {
			parts := strings.Split(errmsg, ":")

			if len(parts) != 2 {
				panic("invalid validation message format")
			}

			key, msg := parts[0], parts[1]

			v.add(key, msg)
		}
	}

	if !v.valid() {
		return v.valerr
	}

	return nil
}
