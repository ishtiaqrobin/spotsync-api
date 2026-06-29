package user

import (
	"github.com/ishtiaqrobin/spotsync-api/internal/auth"
	"github.com/ishtiaqrobin/spotsync-api/internal/config"
	"github.com/ishtiaqrobin/spotsync-api/internal/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// RegisterRoutes wires up all user/auth routes with dependency injection
func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	// Wire dependencies: Repository → Service → Handler
	userRepository := NewRepository(db)
	jwtService := auth.NewJWTService(cfg.JwtSecret)
	userService := NewService(userRepository, jwtService)
	userHandler := NewHandler(userService)

	// Auth routes (public)
	api := e.Group("/api/v1/auth")
	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)

	// Protected example route (optional, for testing middleware)
	api.GET("/me", userHandler.GetMe, middleware.JWTAuth(jwtService))
}
