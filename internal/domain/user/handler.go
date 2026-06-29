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
		return httpresponse.Fail(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}

	if err := c.Validate(&req); err != nil {
		return httpresponse.Fail(c, http.StatusBadRequest, "Validation failed", validation.ParseValidationErrors(err))
	}

	response, err := h.service.Register(req)
	if err != nil {
		if errors.Is(err, ErrEmailAlreadyExists) {
			return httpresponse.Fail(c, http.StatusBadRequest, "Email already registered", nil)
		}
		return httpresponse.Fail(c, http.StatusInternalServerError, "Failed to register user", err.Error())
	}

	return httpresponse.Success(c, http.StatusCreated, "User registered successfully", response)
}

// GetMe returns the current authenticated user's info
func (h *handler) GetMe(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	user, err := h.service.FindByID(userID)
	if err != nil {
		return httpresponse.Fail(c, http.StatusNotFound, "User not found", nil)
	}

	return httpresponse.Success(c, http.StatusOK, "User retrieved successfully", user)
}

// Login handles user login
func (h *handler) Login(c echo.Context) error {
	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {
		return httpresponse.Fail(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}

	if err := c.Validate(&req); err != nil {
		return httpresponse.Fail(c, http.StatusBadRequest, "Validation failed", validation.ParseValidationErrors(err))
	}

	response, err := h.service.Login(req)
	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			return httpresponse.Fail(c, http.StatusUnauthorized, "Invalid email or password", nil)
		}
		return httpresponse.Fail(c, http.StatusInternalServerError, "Something went wrong", err.Error())
	}

	return httpresponse.Success(c, http.StatusOK, "Login successful", response)
}
