package server

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/ishtiaqrobin/spotsync-api/internal/config"
	"github.com/ishtiaqrobin/spotsync-api/internal/domain/reservation"
	"github.com/ishtiaqrobin/spotsync-api/internal/domain/user"
	"github.com/ishtiaqrobin/spotsync-api/internal/domain/zone"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

// CustomValidator wraps go-playground/validator for Echo
type CustomValidator struct {
	validator *validator.Validate
}

// Validate runs struct-level validation
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// Start initializes the Echo server, registers all routes, and starts listening
func Start(db *gorm.DB, cfg *config.Config) {
	// AutoMigrate all domain entities
	if err := db.AutoMigrate(
		&user.User{},
		&zone.ParkingZone{},
		&reservation.Reservation{},
	); err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}
	log.Println("Database migrated successfully")

	e := echo.New()

	// Use custom validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Basic routes
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Welcome to Spotsync API!"})
	})

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})

	// Register domain routes (each domain wires its own DI)
	user.RegisterRoutes(e, db, cfg)
	zone.RegisterRoutes(e, db, cfg)
	reservation.RegisterRoutes(e, db, cfg)

	// Start server
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
