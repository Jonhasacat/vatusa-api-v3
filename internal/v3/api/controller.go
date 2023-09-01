package api

import (
	"github.com/VATUSA/api-v3/internal/database"
	"github.com/VATUSA/api-v3/pkg/auth"
	"github.com/VATUSA/api-v3/pkg/datamodel/response"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetController(c echo.Context) error {
	cidParam := c.Param("cid")
	cid, err := strconv.ParseUint(cidParam, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	controller, err := database.FetchControllerByCID(cid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if auth.CanReadControllerSensitiveData(c) {
		return c.JSON(http.StatusOK, response.MakeController(controller))
	} else if auth.IsAuthenticated(c) {
		return c.JSON(http.StatusOK, response.MakeControllerRedacted(controller))
	} else {
		return c.JSON(http.StatusOK, response.MakeControllerLimited(controller))
	}
}
