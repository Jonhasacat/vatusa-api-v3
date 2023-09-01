package middleware

import (
	db "github.com/VATUSA/api-v3/internal/database"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	ErrorFacilityHeaderRequired     = echo.NewHTTPError(http.StatusBadRequest, "X-Facility header required")
	ErrorTokenHeaderRequired        = echo.NewHTTPError(http.StatusBadRequest, "X-Token header required")
	ErrorInvalidToken               = echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	ErrorDisabledToken              = echo.NewHTTPError(http.StatusForbidden, "token disabled")
	ErrorTokenNotAuthorizedFacility = echo.NewHTTPError(http.StatusForbidden, "token not authorized for facility")
)

func FacilityAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		facilityHeader := c.Request().Header.Get("X-Facility")
		tokenHeader := c.Request().Header.Get("X-Token")
		if facilityHeader == "" {
			return ErrorFacilityHeaderRequired
		}
		if tokenHeader == "" {
			return ErrorTokenHeaderRequired
		}
		apiToken, err := db.FetchAPITokenByToken(tokenHeader)
		if err != nil {
			return ErrorInvalidToken
		}
		if !apiToken.IsEnabled {
			return ErrorDisabledToken
		}
		if facilityHeader != apiToken.APIUser.Facility && apiToken.APIUser.Facility != "*" {
			return ErrorTokenNotAuthorizedFacility
		}
		c.Set("facility", facilityHeader)
		return next(c)
	}
}
