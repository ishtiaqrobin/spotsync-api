package main

import (
	"log"
	"os"

	"github.com/ishtiaqrobin/spotsync-api/internal/config"
	"github.com/ishtiaqrobin/spotsync-api/internal/handler"
	"github.com/ishtiaqrobin/spotsync-api/internal/models"
	"github.com/ishtiaqrobin/spotsync-api/internal/repository"
	"github.com/ishtiaqrobin/spotsync-api/internal/routes"
	"github.com/ishtiaqrobin/spotsync-api/internal/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from system environment")
	}

	// Connect to database
	config.ConnectDatabase()

	// AutoMigrate models
	err := config.DB.AutoMigrate(
		&models.User{},
		&models.ParkingZone{},
		&models.Reservation{},
	)
	if err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}
	log.Println("Database migrated successfully")

	// ---------- Dependency Injection ----------
	// 1. Repository layer
	userRepo := repository.NewUserRepository(config.DB)
	zoneRepo := repository.NewZoneRepository(config.DB)
	reservationRepo := repository.NewReservationRepository(config.DB)

	// 2. Service layer (depends on Repository)
	authService := service.NewAuthService(userRepo)
	zoneService := service.NewZoneService(zoneRepo)
	reservationService := service.NewReservationService(reservationRepo)

	// 3. Handler layer (depends on Service)
	h := &routes.Handlers{
		AuthHandler:        handler.NewAuthHandler(authService),
		ZoneHandler:        handler.NewZoneHandler(zoneService),
		ReservationHandler: handler.NewReservationHandler(reservationService),
	}

	// ---------- Echo setup ----------
	e := echo.New()

	// Basic health check route (temporary, just to confirm server runs)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "SpotSync API is running",
		})
	})

	routes.RegisterRoutes(e, h)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
