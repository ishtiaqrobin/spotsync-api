package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	accessTokenDuration  = 24 * time.Hour
	refreshTokenDuration = 7 * 24 * time.Hour

	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"
)

// JWTClaims represents the claims stored in a JWT token
type JWTClaims struct {
	UserID    uint   `json:"user_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

// JWTService defines the interface for JWT operations
type JWTService interface {
	GenerateAccessToken(userID uint, name, email, role string) (string, error)
	GenerateRefreshToken(userID uint, name, email, role string) (string, error)
	ValidateToken(tokenStr string) (*JWTClaims, error)
}

type jwtService struct {
	secretKey string
}

// NewJWTService creates a new JWTService with the given secret key
func NewJWTService(secretKey string) JWTService {
	if secretKey == "" {
		secretKey = "default_secret_change_me"
	}
	return &jwtService{secretKey: secretKey}
}

// generateToken creates a JWT token with the specified type and duration
func (js *jwtService) generateToken(userID uint, name, email, role, tokenType string, duration time.Duration) (string, error) {
	claims := JWTClaims{
		UserID:    userID,
		Name:      name,
		Email:     email,
		Role:      role,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "spotsync",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(js.secretKey))
}

// GenerateAccessToken creates a short-lived access token
func (js *jwtService) GenerateAccessToken(userID uint, name, email, role string) (string, error) {
	return js.generateToken(userID, name, email, role, TokenTypeAccess, accessTokenDuration)
}

// GenerateRefreshToken creates a long-lived refresh token
func (js *jwtService) GenerateRefreshToken(userID uint, name, email, role string) (string, error) {
	return js.generateToken(userID, name, email, role, TokenTypeRefresh, refreshTokenDuration)
}

// ValidateToken parses and validates a JWT token, returning its claims
func (js *jwtService) ValidateToken(tokenStr string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(js.secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
