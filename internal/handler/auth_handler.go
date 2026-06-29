package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/ishtiaqrobin/spotsync-api/internal/dto"
	"github.com/ishtiaqrobin/spotsync-api/internal/service"
	"github.com/ishtiaqrobin/spotsync-api/internal/utils"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService *service.AuthService
	validate    *validator.Validate
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validate:    validator.New(),
	}
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req dto.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorJSON(c, http.StatusBadRequest, "Invalid request body", err.Error())
	}
	if err := h.validate.Struct(req); err != nil {
		return utils.ErrorJSON(c, http.StatusBadRequest, "Validation failed", err.Error())
	}

	user, err := h.authService.Register(req)
	if err != nil {
		if err == service.ErrEmailAlreadyExists {
			return utils.ErrorJSON(c, http.StatusBadRequest, "Email already registered", err.Error())
		}
		return utils.ErrorJSON(c, http.StatusInternalServerError, "Failed to register user", err.Error())
	}

	res := dto.RegisterResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	return utils.SuccessJSON(c, http.StatusCreated, "User registered successfully", res)
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req dto.LoginRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorJSON(c, http.StatusBadRequest, "Invalid request body", err.Error())
	}
	if err := h.validate.Struct(req); err != nil {
		return utils.ErrorJSON(c, http.StatusBadRequest, "Validation failed", err.Error())
	}

	token, user, err := h.authService.Login(req)
	if err != nil {
		return utils.ErrorJSON(c, http.StatusBadRequest, "Invalid email or password", err.Error())
	}

	res := dto.LoginResponse{
		Token: token,
		User: dto.LoginUserInfo{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
	}

	return utils.SuccessJSON(c, http.StatusOK, "Login successful", res)
}
