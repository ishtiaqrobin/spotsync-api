package routes

import (
	"github.com/ishtiaqrobin/spotsync-api/internal/handler"
	"github.com/ishtiaqrobin/spotsync-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	AuthHandler        *handler.AuthHandler
	ZoneHandler        *handler.ZoneHandler
	ReservationHandler *handler.ReservationHandler
}

func RegisterRoutes(e *echo.Echo, h *Handlers) {
	api := e.Group("/api/v1")

	// ---------- Auth routes (public) ----------
	auth := api.Group("/auth")
	auth.POST("/register", h.AuthHandler.Register)
	auth.POST("/login", h.AuthHandler.Login)

	// ---------- Zone routes ----------
	zones := api.Group("/zones")
	zones.GET("", h.ZoneHandler.GetAllZones)                                              // public
	zones.GET("/:id", h.ZoneHandler.GetZoneByID)                                          // public
	zones.POST("", h.ZoneHandler.CreateZone, middleware.JWTAuth, middleware.RequireAdmin) // admin only

	// ---------- Reservation routes ----------
	reservations := api.Group("/reservations")
	reservations.POST("", h.ReservationHandler.CreateReservation, middleware.JWTAuth)
	reservations.GET("/my-reservations", h.ReservationHandler.GetMyReservations, middleware.JWTAuth)
	reservations.DELETE("/:id", h.ReservationHandler.CancelReservation, middleware.JWTAuth)
	reservations.GET("", h.ReservationHandler.GetAllReservations, middleware.JWTAuth, middleware.RequireAdmin) // admin only
}
