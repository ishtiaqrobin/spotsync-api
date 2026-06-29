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
		return httpresponse.Fail(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}

	if err := c.Validate(&req); err != nil {
		return httpresponse.Fail(c, http.StatusBadRequest, "Validation failed", validation.ParseValidationErrors(err))
	}

	response, err := h.service.CreateReservation(userID, req)
	if err != nil {
		switch {
		case errors.Is(err, ErrZoneFull):
			return httpresponse.Fail(c, http.StatusConflict, "Parking zone is full", nil)
		case errors.Is(err, ErrZoneNotFound):
			return httpresponse.Fail(c, http.StatusNotFound, "Parking zone not found", nil)
		default:
			return httpresponse.Fail(c, http.StatusInternalServerError, "Failed to create reservation", err.Error())
		}
	}

	return httpresponse.Success(c, http.StatusCreated, "Reservation confirmed successfully", response)
}

// GetMyReservations handles retrieving the authenticated user's reservations
func (h *handler) GetMyReservations(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	reservations, err := h.service.GetMyReservations(userID)
	if err != nil {
		return httpresponse.Fail(c, http.StatusInternalServerError, "Failed to fetch reservations", err.Error())
	}

	return httpresponse.Success(c, http.StatusOK, "My reservations retrieved successfully", reservations)
}

// GetAllReservations handles retrieving all reservations (admin only)
func (h *handler) GetAllReservations(c echo.Context) error {
	reservations, err := h.service.GetAllReservations()
	if err != nil {
		return httpresponse.Fail(c, http.StatusInternalServerError, "Failed to fetch reservations", err.Error())
	}

	return httpresponse.Success(c, http.StatusOK, "All reservations retrieved successfully", reservations)
}

// CancelReservation handles cancelling a reservation
func (h *handler) CancelReservation(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	userRole, _ := c.Get("user_role").(string)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return httpresponse.Fail(c, http.StatusBadRequest, "Invalid reservation id", err.Error())
	}

	err = h.service.CancelReservation(uint(id), userID, userRole)
	if err != nil {
		switch {
		case errors.Is(err, ErrReservationNotFound):
			return httpresponse.Fail(c, http.StatusNotFound, "Reservation not found", nil)
		case errors.Is(err, ErrForbidden):
			return httpresponse.Fail(c, http.StatusForbidden, "You can only cancel your own reservations", nil)
		default:
			return httpresponse.Fail(c, http.StatusInternalServerError, "Failed to cancel reservation", err.Error())
		}
	}

	return httpresponse.Success(c, http.StatusOK, "Reservation cancelled successfully", nil)
}
