package facility

import (
	"github.com/VATUSA/api-v3/internal/api/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"net/http"
)

func App() *echo.Echo {
	e := echo.New()
	e.Use(middleware.FacilityAuth)
	e.Use(echomiddleware.CORS())
	e.Use(echomiddleware.LoggerWithConfig(echomiddleware.LoggerConfig{
		Format: "${method} ${uri} - HTTP ${status}\n",
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/roster", GetRoster)
	e.POST("/roster/remove", RemoveFromRoster)

	e.GET("/roster/request", GetPendingRosterRequests)
	e.PUT("/roster/request", ProcessRosterRequest)

	return e
}
