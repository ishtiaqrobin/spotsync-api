package handler

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/ishtiaqrobin/spotsync-api/internal/dto"
	"github.com/ishtiaqrobin/spotsync-api/internal/service"
	"github.com/ishtiaqrobin/spotsync-api/internal/utils"
	"github.com/labstack/echo/v4"
)

type ZoneHandler struct {
	zoneService *service.ZoneService
	validate    *validator.Validate
}

func NewZoneHandler(zoneService *service.ZoneService) *ZoneHandler {
	return &ZoneHandler{
		zoneService: zoneService,
		validate:    validator.New(),
	}
}

func (h *ZoneHandler) CreateZone(c echo.Context) error {
	var req dto.CreateZoneRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorJSON(c, http.StatusBadRequest, "Invalid request body", err.Error())
	}
	if err := h.validate.Struct(req); err != nil {
		return utils.ErrorJSON(c, http.StatusBadRequest, "Validation failed", utils.ParseValidationErrors(err))
	}

	zone, err := h.zoneService.CreateZone(req)
	if err != nil {
		return utils.ErrorJSON(c, http.StatusInternalServerError, "Failed to create parking zone", err.Error())
	}

	res := dto.ZoneResponse{
		ID:             zone.ID,
		Name:           zone.Name,
		Type:           zone.Type,
		TotalCapacity:  zone.TotalCapacity,
		AvailableSpots: zone.TotalCapacity, // new zone, no reservations yet
		PricePerHour:   zone.PricePerHour,
		CreatedAt:      zone.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	return utils.SuccessJSON(c, http.StatusCreated, "Parking zone created successfully", res)
}

func (h *ZoneHandler) GetAllZones(c echo.Context) error {
	zones, err := h.zoneService.GetAllZones()
	if err != nil {
		return utils.ErrorJSON(c, http.StatusInternalServerError, "Failed to fetch parking zones", err.Error())
	}
	return utils.SuccessJSON(c, http.StatusOK, "Parking zones retrieved successfully", zones)
}

func (h *ZoneHandler) GetZoneByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ErrorJSON(c, http.StatusBadRequest, "Invalid zone id", err.Error())
	}

	zone, err := h.zoneService.GetZoneByID(uint(id))
	if err != nil {
		return utils.ErrorJSON(c, http.StatusNotFound, "Parking zone not found", err.Error())
	}

	return utils.SuccessJSON(c, http.StatusOK, "Parking zone retrieved successfully", zone)
}
