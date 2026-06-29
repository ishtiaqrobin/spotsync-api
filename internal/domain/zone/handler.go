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

	response, err := h.service.CreateZone(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create parking zone",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response)
}

// GetAllZones handles retrieving all parking zones
func (h *handler) GetAllZones(c echo.Context) error {
	zones, err := h.service.GetAllZones()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to fetch parking zones",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, zones)
}

// GetZoneByID handles retrieving a single parking zone by ID
func (h *handler) GetZoneByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid zone id",
			Details: err.Error(),
		})
	}

	zone, err := h.service.GetZoneByID(uint(id))
	if err != nil {
		if errors.Is(err, ErrZoneNotFound) {
			return c.JSON(http.StatusNotFound, httpresponse.Error{
				Code:    http.StatusNotFound,
				Message: "Parking zone not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to fetch parking zone",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, zone)
}
