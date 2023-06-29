package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"vatusa-api-v3/auth"
	"vatusa-api-v3/config"
	"vatusa-api-v3/constants"
	"vatusa-api-v3/database"
)

func AuthContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		method := constants.NoAuth
		cookie, err := c.Cookie(config.UserCookieName)
		var controller *database.Controller
		if cookie != nil {
			method = constants.Controller
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
			method = constants.APIUser
			apiToken, err = database.FetchAPITokenByToken(token)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}
			if !apiToken.IsEnabled {
				return echo.NewHTTPError(http.StatusForbidden, "token is disabled")
			}
			apiUser = apiToken.APIUser
		}
		c.Set(constants.FieldMethod, method)
		c.Set(constants.FieldToken, token)
		c.Set(constants.FieldAPIUser, apiUser)
		c.Set(constants.FieldController, controller)
		return next(c)
	}
}
