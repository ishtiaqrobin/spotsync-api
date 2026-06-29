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

	// dsn := os.Getenv("DSN")
	// if dsn == "" {
	// 	// Build DSN from individual DB components (legacy support)
	// 	dsn = buildDSN()
	// }

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

// buildDSN constructs a PostgreSQL DSN string from individual environment variables
// func buildDSN() string {
// 	host := os.Getenv("DB_HOST")
// 	port := os.Getenv("DB_PORT")
// 	user := os.Getenv("DB_USER")
// 	password := os.Getenv("DB_PASSWORD")
// 	dbname := os.Getenv("DB_NAME")
// 	sslmode := os.Getenv("DB_SSLMODE")

// 	if host == "" || user == "" || dbname == "" {
// 		return ""
// 	}
// 	if port == "" {
// 		port = "5432"
// 	}
// 	if sslmode == "" {
// 		sslmode = "require"
// 	}

// 	return fmt.Sprintf(
// 		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
// 		host, port, user, password, dbname, sslmode,
// 	)
// }
