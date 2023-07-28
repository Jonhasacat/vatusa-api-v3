package middleware

import (
	"github.com/VATUSA/api-v3/pkg/auth"
	"github.com/VATUSA/api-v3/pkg/constants"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AuthFacilityATMOrDATM(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		facilityID := constants.Facility(c.Param("facility"))
		if !auth.IsFacilityATMOrDATM(c, facilityID) {
			return echo.NewHTTPError(http.StatusUnauthorized, "not authorized")
		}
		return next(c)
	}
}
