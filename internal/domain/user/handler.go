package user

import (
	"errors"
	"net/http"

	"github.com/ishtiaqrobin/spotsync-api/internal/domain/user/dto"
	"github.com/ishtiaqrobin/spotsync-api/internal/httpresponse"
	"github.com/ishtiaqrobin/spotsync-api/internal/validation"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service *service
}

// NewHandler creates a new user handler
func NewHandler(s *service) *handler {
	return &handler{service: s}
}

// Register handles user registration
func (h *handler) Register(c echo.Context) error {
	var req dto.RegisterRequest

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

	response, err := h.service.Register(req)
	if err != nil {
		if errors.Is(err, ErrEmailAlreadyExists) {
			return c.JSON(http.StatusBadRequest, httpresponse.Error{
				Code:    http.StatusBadRequest,
				Message: "Email already registered",
			})
		}
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to register user",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response)
}

// GetMe returns the current authenticated user's info
func (h *handler) GetMe(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	user, err := h.service.FindByID(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, httpresponse.Error{
			Code:    http.StatusNotFound,
			Message: "User not found",
		})
	}

	return c.JSON(http.StatusOK, user)
}

// Login handles user login
func (h *handler) Login(c echo.Context) error {
	var req dto.LoginRequest

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

	response, err := h.service.Login(req)
	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			return c.JSON(http.StatusUnauthorized, httpresponse.Error{
				Code:    http.StatusUnauthorized,
				Message: "Invalid email or password",
			})
		}
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Something went wrong",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}
