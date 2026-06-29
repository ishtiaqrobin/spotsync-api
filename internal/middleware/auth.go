package middleware

import (
	"net/http"
	"strings"

	"github.com/ishtiaqrobin/spotsync-api/internal/auth"
	"github.com/ishtiaqrobin/spotsync-api/internal/httpresponse"

	"github.com/labstack/echo/v4"
)

// JWTAuth creates a middleware that validates JWT tokens and injects user claims into context
func JWTAuth(jwtService auth.JWTService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, httpresponse.Error{
					Code:    http.StatusUnauthorized,
					Message: "Missing authorization header",
				})
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, httpresponse.Error{
					Code:    http.StatusUnauthorized,
					Message: "Invalid authorization header format",
				})
			}

			claims, err := jwtService.ValidateToken(parts[1])
			if err != nil {
				return c.JSON(http.StatusUnauthorized, httpresponse.Error{
					Code:    http.StatusUnauthorized,
					Message: "Invalid or expired token",
				})
			}
			
			// Reject refresh tokens from being used as access tokens
			if claims.TokenType != auth.TokenTypeAccess {
				return c.JSON(http.StatusUnauthorized, httpresponse.Error{
					Code:    http.StatusUnauthorized,
					Message: "Invalid token type",
				})
			}

			// Inject user claims into Echo context
			c.Set("user_id", claims.UserID)
			c.Set("user_name", claims.Name)
			c.Set("user_email", claims.Email)
			c.Set("user_role", claims.Role)

			return next(c)
		}
	}
}

// RequireAdmin ensures only users with role "admin" can access the route
// Must be used AFTER JWTAuth middleware (depends on c.Get("user_role"))
func RequireAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role, ok := c.Get("user_role").(string)
		if !ok || role != "admin" {
			return c.JSON(http.StatusForbidden, httpresponse.Error{
				Code:    http.StatusForbidden,
				Message: "Admin access required",
			})
		}
		return next(c)
	}
}
