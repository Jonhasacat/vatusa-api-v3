package auth

import (
	db "github.com/VATUSA/api-v3/pkg/database"
	"github.com/VATUSA/api-v3/pkg/facility"
	"github.com/VATUSA/api-v3/pkg/role"
	"github.com/labstack/echo/v4"
)

func IsAuthenticated(c echo.Context) bool {
	return c.Get(FieldMethod) == Controller || c.Get(FieldMethod) == APIUser
}

func IsController(c echo.Context) bool {
	return c.Get(FieldMethod) == Controller
}

func IsStaff(c echo.Context) bool {
	if c.Get(FieldMethod) != Controller {
		return false
	}
	controllerModel := c.Get(FieldController).(db.Controller)
	return role.IsStaff(&controllerModel, controllerModel.Facility)
}

func IsSeniorStaff(c echo.Context) bool {
	if c.Get(FieldMethod) != Controller {
		return false
	}
	controllerModel := c.Get(FieldController).(db.Controller)
	return role.IsSeniorStaff(&controllerModel, controllerModel.Facility)
}
func IsFacilityStaff(c echo.Context, facility facility.Facility) bool {
	if c.Get(FieldMethod) != Controller {
		return false
	}
	controllerModel := c.Get(FieldController).(db.Controller)
	return role.IsStaff(&controllerModel, facility)
}
func IsFacilityTrainingStaff(c echo.Context, facility facility.Facility) bool {
	if c.Get(FieldMethod) != Controller {
		return false
	}
	controllerModel := c.Get(FieldController).(db.Controller)
	return role.IsTrainingStaff(&controllerModel, facility)
}

func IsFacilityInstructor(c echo.Context, facility facility.Facility) bool {
	if c.Get(FieldMethod) != Controller {
		return false
	}
	controllerModel := c.Get(FieldController).(db.Controller)
	return role.IsInstructor(&controllerModel, facility)
}

func IsFacilitySeniorStaff(c echo.Context, facility facility.Facility) bool {
	if c.Get(FieldMethod) != Controller {
		return false
	}
	controllerModel := c.Get(FieldController).(db.Controller)
	return role.IsSeniorStaff(&controllerModel, facility)
}

func IsFacilityATMOrDATM(c echo.Context, facility facility.Facility) bool {
	if c.Get(FieldMethod) != Controller {
		return false
	}
	controllerModel := c.Get(FieldController).(db.Controller)
	return role.IsATMOrDATM(&controllerModel, facility)
}

func IsFacilityToken(c echo.Context, facility facility.Facility) bool {
	if c.Get(FieldMethod) != APIUser {
		return false
	}
	apiUser := c.Get(FieldAPIUser).(db.APIUser)
	return apiUser.Facility == facility || apiUser.Facility == "*"
}

func CanReadControllerSensitiveData(c echo.Context) bool {
	if c.Get(FieldMethod) == NoAuth {
		return false
	} else if c.Get(FieldMethod) == Controller {
		return IsSeniorStaff(c)
	} else if c.Get(FieldMethod) == APIUser {
		return true
	}
	return false
}
