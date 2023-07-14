package middleware

import (
	"github.com/VATUSA/api-v3/pkg/auth"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AuthControllerOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !auth.IsController(c) {
			return echo.NewHTTPError(http.StatusUnauthorized, "not authorized")
		}
		return next(c)
	}
}
