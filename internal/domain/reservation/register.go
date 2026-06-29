package reservation

import (
	"github.com/ishtiaqrobin/spotsync-api/internal/auth"
	"github.com/ishtiaqrobin/spotsync-api/internal/config"
	"github.com/ishtiaqrobin/spotsync-api/internal/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// RegisterRoutes wires up all reservation routes with dependency injection
func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	// Wire dependencies: Repository → Service → Handler
	reservationRepository := NewRepository(db)
	reservationService := NewService(reservationRepository)
	reservationHandler := NewHandler(reservationService)

	jwtService := auth.NewJWTService(cfg.JwtSecret)

	// Reservation routes (all require authentication)
	api := e.Group("/api/v1/reservations", middleware.JWTAuth(jwtService))

	// Specific routes first (before parameterized routes)
	api.POST("", reservationHandler.CreateReservation)
	api.GET("/my-reservations", reservationHandler.GetMyReservations)
	api.GET("", reservationHandler.GetAllReservations, middleware.RequireAdmin)
	api.DELETE("/:id", reservationHandler.CancelReservation)
}
