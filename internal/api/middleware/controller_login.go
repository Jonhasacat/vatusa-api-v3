package middleware

import (
	"github.com/VATUSA/api-v3/internal/config"
	"github.com/VATUSA/api-v3/pkg/auth"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	ErrorBadToken = echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
)

func ControllerAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, _ := c.Cookie(config.UserCookieName)
		if cookie != nil {
			controller, err := auth.GetControllerForJWT(cookie.Value)
			if err != nil {
				return ErrorBadToken
			}
			c.Set(auth.FieldController, controller)
		}
		return next(c)
	}
}
