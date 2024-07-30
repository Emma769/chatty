package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, code int, value any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(value)
}

func WriteJsonE(w http.ResponseWriter, code int, value any) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(value)
}

func ReadJson(w http.ResponseWriter, r *http.Request, value any) error {
	return nil
}
