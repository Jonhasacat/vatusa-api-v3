package api

import (
	"github.com/VATUSA/api-v3/pkg/auth"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetMyInfo(c echo.Context) error {
	userId, err := auth.GetRequestUserID(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "User token not present or malformed")
	}
	if userId == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "TODO")
	}
	return c.JSON(http.StatusOK, "")
}
