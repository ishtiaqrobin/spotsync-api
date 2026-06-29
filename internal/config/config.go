package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all application configuration
type Config struct {
	Port      string
	Dsn       string
	JwtSecret string
}

// LoadEnv loads environment variables from .env file (if present)
// and returns a Config struct. In production, env vars come from the platform.
func LoadEnv() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	cfg := &Config{
		Port: os.Getenv("PORT"),
		// Dsn:       dsn,
		Dsn:       os.Getenv("DSN"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}

	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	if cfg.Dsn == "" {
		log.Fatal("DSN environment variable is required")
	}
	if cfg.JwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}

	return cfg
}
