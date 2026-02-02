package config

import (
	"os"
)

type Config struct {
	Port string
	DB   struct {
		DSN string
	}
	JWT struct {
		Secret string
	}
}

func Load() Config {
	var cfg Config

	cfg.Port = getEnv("PORT", "4000")
	cfg.DB.DSN = getEnv(
		"DB_DSN",
		"postgres://postgres:chintukr1904@@localhost:5432/studentdb?sslmode=disable",
	)
	cfg.JWT.Secret = getEnv("JWT_SECRET", "dev-secret-change=me")
	return cfg
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
