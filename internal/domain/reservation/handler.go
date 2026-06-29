package reservation

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ishtiaqrobin/spotsync-api/internal/domain/reservation/dto"
	"github.com/ishtiaqrobin/spotsync-api/internal/httpresponse"
	"github.com/ishtiaqrobin/spotsync-api/internal/validation"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service *service
}

// NewHandler creates a new reservation handler
func NewHandler(s *service) *handler {
	return &handler{service: s}
}

// CreateReservation handles creating a new reservation
func (h *handler) CreateReservation(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	var req dto.CreateReservationRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Details: validation.ParseValidationErrors(err),
		})
	}

	response, err := h.service.CreateReservation(userID, req)
	if err != nil {
		switch {
		case errors.Is(err, ErrZoneFull):
			return c.JSON(http.StatusConflict, httpresponse.Error{
				Code:    http.StatusConflict,
				Message: "Parking zone is full",
			})
		case errors.Is(err, ErrZoneNotFound):
			return c.JSON(http.StatusNotFound, httpresponse.Error{
				Code:    http.StatusNotFound,
				Message: "Parking zone not found",
			})
		default:
			return c.JSON(http.StatusInternalServerError, httpresponse.Error{
				Code:    http.StatusInternalServerError,
				Message: "Failed to create reservation",
				Details: err.Error(),
			})
		}
	}

	return c.JSON(http.StatusCreated, response)
}

// GetMyReservations handles retrieving the authenticated user's reservations
func (h *handler) GetMyReservations(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	reservations, err := h.service.GetMyReservations(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to fetch reservations",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, reservations)
}

// GetAllReservations handles retrieving all reservations (admin only)
func (h *handler) GetAllReservations(c echo.Context) error {
	reservations, err := h.service.GetAllReservations()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to fetch reservations",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, reservations)
}

// CancelReservation handles cancelling a reservation
func (h *handler) CancelReservation(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	userRole, _ := c.Get("user_role").(string)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid reservation id",
			Details: err.Error(),
		})
	}

	err = h.service.CancelReservation(uint(id), userID, userRole)
	if err != nil {
		switch {
		case errors.Is(err, ErrReservationNotFound):
			return c.JSON(http.StatusNotFound, httpresponse.Error{
				Code:    http.StatusNotFound,
				Message: "Reservation not found",
			})
		case errors.Is(err, ErrForbidden):
			return c.JSON(http.StatusForbidden, httpresponse.Error{
				Code:    http.StatusForbidden,
				Message: "You can only cancel your own reservations",
			})
		default:
			return c.JSON(http.StatusInternalServerError, httpresponse.Error{
				Code:    http.StatusInternalServerError,
				Message: "Failed to cancel reservation",
				Details: err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Reservation cancelled successfully",
	})
}
