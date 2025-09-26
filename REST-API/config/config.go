package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all environment-driven settings.
type Config struct {
	ServerPort  string
	DatabaseURL string
}

// New loads variables and returns a Config instance.
// It first tries a .env file (for local dev) but works
// fine in Docker/production where env vars are already set.
func New() *Config {
	// Load .env only if present—no error if running in Docker
	_ = godotenv.Load()

	c := &Config{
		ServerPort:  getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
	}
	if c.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}
	return c
}

// helper for default values
func getEnv(key, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return def
}
