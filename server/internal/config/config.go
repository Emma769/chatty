package config

import (
	"strconv"
	"syscall"
)

type Getter struct{}

func (g Getter) Get(key string, fallback string) string {
	value, ok := syscall.Getenv(key)

	if !ok {
		return fallback
	}

	return value
}

func (g Getter) GetInt(key string, fallback int) int {
	value, ok := syscall.Getenv(key)

	if !ok {
		return fallback
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return i
}
