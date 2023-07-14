package middleware

import (
	"fmt"
	"github.com/VATUSA/api-v3/pkg/auth"
	"github.com/VATUSA/api-v3/pkg/facility"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AuthFacilityStaffOrToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		facilityID := facility.Facility(c.Param("facility"))
		if !auth.IsFacilityStaff(c, facilityID) && !auth.IsFacilityToken(c, facilityID) {
			return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("Not authorized for %s", facilityID))
		}
		return next(c)
	}
}
