package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, code int, value any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(value)
}

func WriteJsonE(w http.ResponseWriter, code int, value any) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(value)
}

func ReadJSON(w http.ResponseWriter, r *http.Request, value any) error {
	defer func() { _ = r.Body.Close() }()

	body := http.MaxBytesReader(w, r.Body, 1_048_576)
	err := json.NewDecoder(body).Decode(&value)

	var syntaxErr *json.SyntaxError

	if err != nil && errors.As(err, &syntaxErr) {
		return fmt.Errorf("invalid json at %d", syntaxErr.Offset)
	}

	var typeErr *json.UnmarshalTypeError

	if err != nil && errors.As(err, &typeErr) {
		if typeErr.Field != "" {
			return fmt.Errorf("invalid json at %s", typeErr.Field)
		}
		return fmt.Errorf("invalid json at %d", typeErr.Offset)
	}

	if err != nil && errors.Is(err, io.EOF) {
		return errors.New("request body has no content")
	}

	if err != nil && errors.Is(err, io.ErrUnexpectedEOF) {
		return errors.New("malformed json")
	}

	return err
}

type err interface {
	error
	Code() int
}

type handlerfn func(http.ResponseWriter, *http.Request) error

func Wrap[E err](fn handlerfn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e E

		err := fn(w, r)

		if err != nil && errors.As(err, &e) {
			_ = WriteJsonE(w, e.Code(), map[string]string{"detail": e.Error()})
			return
		}

		if err != nil {
			panic(err)
		}
	}
}
