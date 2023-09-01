package middleware

import (
	"github.com/VATUSA/api-v3/internal/config"
	"github.com/VATUSA/api-v3/internal/database"
	"github.com/VATUSA/api-v3/pkg/auth"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func AuthContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		method := auth.NoAuth
		cookie, err := c.Cookie(config.UserCookieName)
		var controller *database.Controller
		if cookie != nil {
			method = auth.Controller
			controller, err = auth.GetControllerForJWT(cookie.Value)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid JWT")
			}
		}
		header := c.Request().Header.Get("Authorization")
		var token string
		var apiToken *database.APIToken
		var apiUser *database.APIUser
		if header != "" {
			fields := strings.Fields(header)
			if len(fields) > 2 {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid Token format")
			}
			if fields[0] != "Token" {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token scheme")
			}
			token = fields[1]
			method = auth.APIUser
			apiToken, err = database.FetchAPITokenByToken(token)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}
			if !apiToken.IsEnabled {
				return echo.NewHTTPError(http.StatusForbidden, "token is disabled")
			}
			apiUser = apiToken.APIUser
		}
		c.Set(auth.FieldMethod, method)
		c.Set(auth.FieldToken, token)
		c.Set(auth.FieldAPIUser, apiUser)
		c.Set(auth.FieldController, controller)
		return next(c)
	}
}
