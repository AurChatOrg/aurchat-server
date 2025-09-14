package config

import (
	"strconv"
	"sync"
)

var (
	once sync.Once
	cfg  *Config
)

// Load Load env
func Load() *Config {
	once.Do(func() {
		cfg = &Config{
			App: App{
				Name: readEnv("APP_NAME", "aurchat-gateway"),
				Env:  readEnv("APP_ENV", "dev"),
			},
			HTTP: HTTP{
				Listen: readEnv("HTTP_LISTEN", ":8080"),
			},
			Redis: Redis{
				Addr: readEnv("REDIS_ADDR", "127.0.0.1:6379"),
				DB:   mustAtoi(readEnv("REDIS_DB", "0")),
			},
			NATS: NATS{
				URL: readEnv("NATS_URL", "nats://127.0.0.1:4222"),
			},
			DSN: DSN{
				Postgres: readEnv("POSTGRES_DSN", "host=localhost user=user password=pass dbname=aurchat port=5432 sslmode=disable"),
			},
		}
	})
	return cfg
}

func mustAtoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
