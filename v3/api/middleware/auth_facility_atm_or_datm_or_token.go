package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"vatusa-api-v3/auth"
)

func AuthFacilityATMOrDATMOrToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		facility := c.Param("facility")
		if !auth.IsFacilityATMOrDATM(c, facility) && !auth.IsFacilityToken(c, facility) {
			return echo.NewHTTPError(http.StatusUnauthorized, "not authorized")
		}
		return next(c)
	}
}
