package config

import (
	"strconv"
	"sync"
	"syscall"
)

type Config struct {
	PORT                  int
	POSTGRES_URI          string
	DB_MAX_OPEN_CONNS     int
	DB_MAX_IDLE_CONNS     int
	DB_CONN_MAX_IDLE_TIME int
	READ_TIMEOUT          int
	WRITE_TIMEOUT         int
}

var (
	once sync.Once
	cfg  *Config
)

func Load() *Config {
	once.Do(func() {
		cfg = &Config{
			PORT:                  getInt("PORT", 4000),
			POSTGRES_URI:          get("POSTGRES_URI", ""),
			DB_MAX_OPEN_CONNS:     getInt("DB_MAX_OPEN_CONNS", 25),
			DB_MAX_IDLE_CONNS:     getInt("DB_MAX_IDLE_CONNS", 25),
			DB_CONN_MAX_IDLE_TIME: getInt("DB_CONN_MAX_IDLE_TIME", 15),
			READ_TIMEOUT:          getInt("READ_TIMEOUT", 15),
			WRITE_TIMEOUT:         getInt("WRITE_TIMEOUT", 15),
		}
	})

	return cfg
}

func get(key string, fallback string) string {
	value, ok := syscall.Getenv(key)

	if !ok {
		return fallback
	}

	return value
}

func getInt(key string, fallback int) int {
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
