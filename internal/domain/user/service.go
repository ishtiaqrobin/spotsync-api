package user

import (
	"fmt"

	"github.com/ishtiaqrobin/spotsync-api/internal/auth"
	"github.com/ishtiaqrobin/spotsync-api/internal/domain/user/dto"
)

// ErrInvalidCredentials is returned when login credentials are incorrect
var ErrInvalidCredentials = fmt.Errorf("invalid email or password")

type service struct {
	repo       Repository
	jwtService auth.JWTService
}

// NewService creates a new user service with the given repository and JWT service
func NewService(repo Repository, jwtService auth.JWTService) *service {
	return &service{repo: repo, jwtService: jwtService}
}

// Register creates a new user after hashing the password
func (s *service) Register(req dto.RegisterRequest) (*dto.Response, error) {
	// Check for duplicate email
	existing, _ := s.repo.FindByEmail(req.Email)
	if existing != nil {
		return nil, ErrEmailAlreadyExists
	}

	user := User{
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}

	// Default role to "driver" if not specified
	if user.Role == "" {
		user.Role = "driver"
	}

	// Hash password
	if err := user.hashPassword(req.Password); err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Save to database
	if err := s.repo.Create(&user); err != nil {
		return nil, err
	}

	return user.ToResponse(), nil
}

// FindByID retrieves a user by their ID
func (s *service) FindByID(id uint) (*dto.Response, error) {
	user, err := s.repo.FindByID(id)
	if err != nil || user == nil {
		return nil, fmt.Errorf("user not found")
	}
	return user.ToResponse(), nil
}

// Login verifies credentials and returns a JWT token
func (s *service) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil || user == nil {
		return nil, ErrInvalidCredentials
	}

	if err := user.checkPassword(req.Password); err != nil {
		return nil, ErrInvalidCredentials
	}

	// Generate JWT token
	token, err := s.jwtService.GenerateAccessToken(user.ID, user.Name, user.Email, user.Role)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &dto.LoginResponse{
		Token: token,
		User: dto.LoginUserInfo{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
	}, nil
}
