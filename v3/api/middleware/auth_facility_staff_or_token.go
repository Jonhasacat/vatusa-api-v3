package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"vatusa-api-v3/auth"
)

func AuthFacilityStaffOrToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		facility := c.Param("facility")
		if !auth.IsFacilityStaff(c, facility) && !auth.IsFacilityToken(c, facility) {
			return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("Not authorized for %s", facility))
		}
		return next(c)
	}
}