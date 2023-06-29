package auth

import (
	"github.com/labstack/echo/v4"
	"vatusa-api-v3/constants"
	"vatusa-api-v3/database"
)

func IsAuthenticated(c echo.Context) bool {
	return c.Get(constants.FieldMethod) == constants.Controller || c.Get(constants.FieldMethod) == constants.APIUser
}

func IsController(c echo.Context) bool {
	return c.Get(constants.FieldMethod) == constants.Controller
}

func IsStaff(c echo.Context) bool {
	if c.Get(constants.FieldMethod) != constants.Controller {
		return false
	}
	controller := c.Get(constants.FieldController).(database.Controller)
	return controller.IsStaff(controller.Facility)
}

func IsSeniorStaff(c echo.Context) bool {
	if c.Get(constants.FieldMethod) != constants.Controller {
		return false
	}
	controller := c.Get(constants.FieldController).(database.Controller)
	return controller.IsSeniorStaff(controller.Facility)
}
func IsFacilityStaff(c echo.Context, facility string) bool {
	if c.Get(constants.FieldMethod) != constants.Controller {
		return false
	}
	controller := c.Get(constants.FieldController).(database.Controller)
	return controller.IsStaff(facility)
}
func IsFacilityTrainingStaff(c echo.Context, facility string) bool {
	if c.Get(constants.FieldMethod) != constants.Controller {
		return false
	}
	controller := c.Get(constants.FieldController).(database.Controller)
	return controller.IsTrainingStaff(facility)
}

func IsFacilityInstructor(c echo.Context, facility string) bool {
	if c.Get(constants.FieldMethod) != constants.Controller {
		return false
	}
	controller := c.Get(constants.FieldController).(database.Controller)
	return controller.IsInstructor(facility)
}

func IsFacilitySeniorStaff(c echo.Context, facility string) bool {
	if c.Get(constants.FieldMethod) != constants.Controller {
		return false
	}
	controller := c.Get(constants.FieldController).(database.Controller)
	return controller.IsSeniorStaff(facility)
}

func IsFacilityATMOrDATM(c echo.Context, facility string) bool {
	if c.Get(constants.FieldMethod) != constants.Controller {
		return false
	}
	controller := c.Get(constants.FieldController).(database.Controller)
	return controller.IsATMOrDATM(facility)
}

func IsFacilityToken(c echo.Context, facility string) bool {
	if c.Get(constants.FieldMethod) != constants.APIUser {
		return false
	}
	apiUser := c.Get(constants.FieldAPIUser).(database.APIUser)
	return apiUser.Facility == facility || apiUser.Facility == "*"
}

func CanReadControllerSensitiveData(c echo.Context) bool {
	if c.Get(constants.FieldMethod) == constants.NoAuth {
		return false
	} else if c.Get(constants.FieldMethod) == constants.Controller {
		return IsSeniorStaff(c)
	} else if c.Get(constants.FieldMethod) == constants.APIUser {
		return true
	}
	return false
}
