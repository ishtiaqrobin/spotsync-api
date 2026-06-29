package handler

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/ishtiaqrobin/spotsync-api/internal/dto"
	"github.com/ishtiaqrobin/spotsync-api/internal/repository"
	"github.com/ishtiaqrobin/spotsync-api/internal/service"
	"github.com/ishtiaqrobin/spotsync-api/internal/utils"
	"github.com/labstack/echo/v4"
)

type ReservationHandler struct {
	reservationService *service.ReservationService
	validate           *validator.Validate
}

func NewReservationHandler(reservationService *service.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		reservationService: reservationService,
		validate:           validator.New(),
	}
}

func (h *ReservationHandler) CreateReservation(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	var req dto.CreateReservationRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorJSON(c, http.StatusBadRequest, "Invalid request body", err.Error())
	}
	if err := h.validate.Struct(req); err != nil {
		return utils.ErrorJSON(c, http.StatusBadRequest, "Validation failed", utils.ParseValidationErrors(err))
	}

	reservation, err := h.reservationService.CreateReservation(userID, req)
	if err != nil {
		switch err {
		case repository.ErrZoneFull:
			return utils.ErrorJSON(c, http.StatusConflict, "Parking zone is full", err.Error())
		case repository.ErrZoneNotFound:
			return utils.ErrorJSON(c, http.StatusNotFound, "Parking zone not found", err.Error())
		default:
			return utils.ErrorJSON(c, http.StatusInternalServerError, "Failed to create reservation", err.Error())
		}
	}

	res := dto.ReservationResponse{
		ID:           reservation.ID,
		UserID:       reservation.UserID,
		ZoneID:       reservation.ZoneID,
		LicensePlate: reservation.LicensePlate,
		Status:       reservation.Status,
		CreatedAt:    reservation.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:    reservation.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	return utils.SuccessJSON(c, http.StatusCreated, "Reservation confirmed successfully", res)
}

func (h *ReservationHandler) GetMyReservations(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	reservations, err := h.reservationService.GetMyReservations(userID)
	if err != nil {
		return utils.ErrorJSON(c, http.StatusInternalServerError, "Failed to fetch reservations", err.Error())
	}

	var result []dto.MyReservationResponse
	for _, res := range reservations {
		result = append(result, dto.MyReservationResponse{
			ID:           res.ID,
			LicensePlate: res.LicensePlate,
			Status:       res.Status,
			Zone: dto.ReservationZoneInfo{
				ID:   res.Zone.ID,
				Name: res.Zone.Name,
				Type: res.Zone.Type,
			},
			CreatedAt: res.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return utils.SuccessJSON(c, http.StatusOK, "My reservations retrieved successfully", result)
}

func (h *ReservationHandler) GetAllReservations(c echo.Context) error {
	reservations, err := h.reservationService.GetAllReservations()
	if err != nil {
		return utils.ErrorJSON(c, http.StatusInternalServerError, "Failed to fetch reservations", err.Error())
	}
	return utils.SuccessJSON(c, http.StatusOK, "All reservations retrieved successfully", reservations)
}

func (h *ReservationHandler) CancelReservation(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	userRole, _ := c.Get("user_role").(string)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ErrorJSON(c, http.StatusBadRequest, "Invalid reservation id", err.Error())
	}

	err = h.reservationService.CancelReservation(uint(id), userID, userRole)
	if err != nil {
		switch err {
		case service.ErrForbidden:
			return utils.ErrorJSON(c, http.StatusForbidden, "You can only cancel your own reservation", err.Error())
		default:
			return utils.ErrorJSON(c, http.StatusNotFound, "Reservation not found", err.Error())
		}
	}

	return utils.SuccessJSON(c, http.StatusOK, "Reservation cancelled successfully", nil)
}
