package config

import (
	"cmp"
	"os"
	"strconv"
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

func New() *Config {
	return &Config{
		PORT:                  cmp.Or(getint("PORT"), 9000),
		POSTGRES_URI:          get("POSTGRES_URI"),
		DB_MAX_OPEN_CONNS:     cmp.Or(getint("DB_MAX_OPEN_CONNS"), 25),
		DB_MAX_IDLE_CONNS:     cmp.Or(getint("DB_MAX_IDLE_CONNS"), 25),
		DB_CONN_MAX_IDLE_TIME: cmp.Or(getint("DB_CONN_MAX_IDLE_TIME"), 15),
		READ_TIMEOUT:          cmp.Or(getint("READ_TIMEOUT"), 15),
		WRITE_TIMEOUT:         cmp.Or(getint("WRITE_TIMEOUT"), 15),
	}
}

func get(k string) string {
	return os.Getenv(k)
}

func getint(k string) int {
	i, _ := strconv.Atoi(os.Getenv(k))
	return i
}
