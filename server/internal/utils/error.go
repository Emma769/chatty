package utils

import (
	"errors"
	"net/http"
)

var ErrUnimplemented = errors.New("unimplemented")

type handlerError struct {
	code   int
	errmsg string
}

func (e handlerError) Error() string {
	return e.errmsg
}

func NewError(code int, errmsg string) *handlerError {
	return &handlerError{code, errmsg}
}

var (
	ErrBadRequest   = NewError(http.StatusBadRequest, "bad request")
	ErrNotFound     = NewError(http.StatusNotFound, "not found")
	ErrUnauthorized = NewError(http.StatusUnauthorized, "unauthorized")
	ErrForbidden    = NewError(http.StatusForbidden, "forbidden")
)

type handlerfn func(http.ResponseWriter, *http.Request) error

func Wrap(fn handlerfn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)

		var e *handlerError

		if err != nil && errors.As(err, &e) {
			_ = WriteJsonE(w, e.code, map[string]string{"detail": e.errmsg})
			return
		}

		if err != nil {
			panic(err)
		}
	}
}
