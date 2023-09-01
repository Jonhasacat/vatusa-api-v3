package auth

import (
	db "github.com/VATUSA/api-v3/internal/database"
	"github.com/labstack/echo/v4"
)

func RequestController(c echo.Context) db.Controller {
	return c.Get(FieldController).(db.Controller)
}
