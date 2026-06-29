package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ishtiaqrobin/spotsync-api/internal/utils"
	"github.com/labstack/echo/v4"
)

// JWTAuth verifies the Bearer token and injects user id + role into context
func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return utils.ErrorJSON(c, http.StatusUnauthorized, "Missing authorization token", nil)
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return utils.ErrorJSON(c, http.StatusUnauthorized, "Invalid authorization header format", nil)
		}

		tokenString := parts[1]
		secret := os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			return utils.ErrorJSON(c, http.StatusUnauthorized, "Invalid or expired token", nil)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return utils.ErrorJSON(c, http.StatusUnauthorized, "Invalid token claims", nil)
		}

		// Inject claims into Echo context for handlers to use
		idFloat, ok := claims["id"].(float64)
		if !ok {
			return utils.ErrorJSON(c, http.StatusUnauthorized, "Invalid token payload", nil)
		}

		c.Set("user_id", uint(idFloat))
		c.Set("user_role", claims["role"])

		return next(c)
	}
}
