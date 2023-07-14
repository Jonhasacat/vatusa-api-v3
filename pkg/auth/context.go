package auth

import (
	"github.com/VATUSA/api-v3/pkg/controller"
	db "github.com/VATUSA/api-v3/pkg/database"
	"github.com/VATUSA/api-v3/pkg/facility"
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
	return controller.IsStaff(&controllerModel, controllerModel.Facility)
}

func IsSeniorStaff(c echo.Context) bool {
	if c.Get(FieldMethod) != Controller {
		return false
	}
	controllerModel := c.Get(FieldController).(db.Controller)
	return controller.IsSeniorStaff(&controllerModel, controllerModel.Facility)
}
func IsFacilityStaff(c echo.Context, facility facility.Facility) bool {
	if c.Get(FieldMethod) != Controller {
		return false
	}
	controllerModel := c.Get(FieldController).(db.Controller)
	return controller.IsStaff(&controllerModel, facility)
}
func IsFacilityTrainingStaff(c echo.Context, facility facility.Facility) bool {
	if c.Get(FieldMethod) != Controller {
		return false
	}
	controllerModel := c.Get(FieldController).(db.Controller)
	return controller.IsTrainingStaff(&controllerModel, facility)
}

func IsFacilityInstructor(c echo.Context, facility facility.Facility) bool {
	if c.Get(FieldMethod) != Controller {
		return false
	}
	controllerModel := c.Get(FieldController).(db.Controller)
	return controller.IsInstructor(&controllerModel, facility)
}

func IsFacilitySeniorStaff(c echo.Context, facility facility.Facility) bool {
	if c.Get(FieldMethod) != Controller {
		return false
	}
	controllerModel := c.Get(FieldController).(db.Controller)
	return controller.IsSeniorStaff(&controllerModel, facility)
}

func IsFacilityATMOrDATM(c echo.Context, facility facility.Facility) bool {
	if c.Get(FieldMethod) != Controller {
		return false
	}
	controllerModel := c.Get(FieldController).(db.Controller)
	return controller.IsATMOrDATM(&controllerModel, facility)
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
