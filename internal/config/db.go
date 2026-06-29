package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectDatabase establishes a connection to PostgreSQL using the DSN from Config.
func ConnectDatabase(cfg *Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Database connection successful")
	return db
}
