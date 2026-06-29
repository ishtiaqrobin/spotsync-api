package zone

import (
	"github.com/ishtiaqrobin/spotsync-api/internal/auth"
	"github.com/ishtiaqrobin/spotsync-api/internal/config"
	"github.com/ishtiaqrobin/spotsync-api/internal/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// RegisterRoutes wires up all parking zone routes with dependency injection
func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	// Wire dependencies: Repository → Service → Handler
	zoneRepository := NewRepository(db)
	zoneService := NewService(zoneRepository)
	zoneHandler := NewHandler(zoneService)

	jwtService := auth.NewJWTService(cfg.JwtSecret)

	// Zone routes
	api := e.Group("/api/v1/zones")

	// Public routes
	api.GET("", zoneHandler.GetAllZones)
	api.GET("/:id", zoneHandler.GetZoneByID)

	// Admin-only routes
	api.POST("", zoneHandler.CreateZone, middleware.JWTAuth(jwtService), middleware.RequireAdmin)
}
