package zone

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ishtiaqrobin/spotsync-api/internal/domain/zone/dto"
	"github.com/ishtiaqrobin/spotsync-api/internal/httpresponse"
	"github.com/ishtiaqrobin/spotsync-api/internal/validation"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service *service
}

// NewHandler creates a new parking zone handler
func NewHandler(s *service) *handler {
	return &handler{service: s}
}

// CreateZone handles creating a new parking zone (admin only)
func (h *handler) CreateZone(c echo.Context) error {
	var req dto.CreateZoneRequest

	if err := c.Bind(&req); err != nil {
		return httpresponse.Fail(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}

	if err := c.Validate(&req); err != nil {
		return httpresponse.Fail(c, http.StatusBadRequest, "Validation failed", validation.ParseValidationErrors(err))
	}

	response, err := h.service.CreateZone(req)
	if err != nil {
		return httpresponse.Fail(c, http.StatusInternalServerError, "Failed to create parking zone", err.Error())
	}

	return httpresponse.Success(c, http.StatusCreated, "Parking zone created successfully", response)
}

// GetAllZones handles retrieving all parking zones
func (h *handler) GetAllZones(c echo.Context) error {
	zones, err := h.service.GetAllZones()
	if err != nil {
		return httpresponse.Fail(c, http.StatusInternalServerError, "Failed to fetch parking zones", err.Error())
	}

	return httpresponse.Success(c, http.StatusOK, "Parking zones retrieved successfully", zones)
}

// GetZoneByID handles retrieving a single parking zone by ID
func (h *handler) GetZoneByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return httpresponse.Fail(c, http.StatusBadRequest, "Invalid zone id", err.Error())
	}

	zone, err := h.service.GetZoneByID(uint(id))
	if err != nil {
		if errors.Is(err, ErrZoneNotFound) {
			return httpresponse.Fail(c, http.StatusNotFound, "Parking zone not found", nil)
		}
		return httpresponse.Fail(c, http.StatusInternalServerError, "Failed to fetch parking zone", err.Error())
	}

	return httpresponse.Success(c, http.StatusOK, "Parking zone retrieved successfully", zone)
}
