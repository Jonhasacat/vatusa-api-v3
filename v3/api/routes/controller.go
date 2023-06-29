package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"vatusa-api-v3/auth"
	"vatusa-api-v3/database"
	"vatusa-api-v3/v3/model"
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
		return c.JSON(http.StatusOK, model.MakeController(controller))
	} else if auth.IsAuthenticated(c) {
		return c.JSON(http.StatusOK, model.MakeControllerRedacted(controller))
	} else {
		return c.JSON(http.StatusOK, model.MakeControllerLimited(controller))
	}
}
