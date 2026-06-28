package main

import (
	"log"
	"os"

	"github.com/ishtiaqrobin/spotsync-api/internal/config"
	"github.com/ishtiaqrobin/spotsync-api/internal/models"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from system environment")
	}

	// Connect to database
	config.ConnectDatabase()

	// AutoMigrate models -> creates/updates tables based on structs
	err := config.DB.AutoMigrate(
		&models.User{},
		&models.ParkingZone{},
		&models.Reservation{},
	)
	if err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}
	log.Println("Database migrated successfully")

	// Initialize Echo instance
	e := echo.New()

	// Basic health check route (temporary, just to confirm server runs)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "SpotSync API is running",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
