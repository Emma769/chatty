package config

import (
	"os"
	"strconv"
	"time"
)

type Getter struct{}

func (Getter) GetString(k string) string {
	return os.Getenv(k)
}

func (Getter) GetInt(k string) int {
	val, _ := strconv.Atoi(os.Getenv(k))
	return val
}

func (Getter) GetDuration(k string) time.Duration {
	val, _ := time.ParseDuration(os.Getenv(k))
	return val
}
