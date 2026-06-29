package middleware

import (
	"net/http"

	"github.com/ishtiaqrobin/spotsync-api/internal/utils"
	"github.com/labstack/echo/v4"
)

// RequireAdmin ensures only users with role "admin" can access the route
// Must be used AFTER JWTAuth middleware (depends on c.Get("user_role"))
func RequireAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role, ok := c.Get("user_role").(string)
		if !ok || role != "admin" {
			return utils.ErrorJSON(c, http.StatusForbidden, "Admin access required", nil)
		}
		return next(c)
	}
}
