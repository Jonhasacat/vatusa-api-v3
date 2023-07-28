package core

import (
	"github.com/VATUSA/api-v3/pkg/constants"
	db "github.com/VATUSA/api-v3/pkg/database"
)

func HasRole(c *db.Controller, role constants.Role, facility constants.Facility) bool {
	for _, r := range c.Roles {
		if r.Role == role && r.Facility == facility {
			return true
		}
	}
	return false
}

func IsStaff(c *db.Controller, facility constants.Facility) bool {
	return IsSeniorStaff(c, facility) ||
		HasRole(c, constants.WebMaster, facility) ||
		HasRole(c, constants.FacilityEngineer, facility) ||
		HasRole(c, constants.EventCoordinator, facility) ||
		IsTrainingStaff(c, facility)
}

func IsSeniorStaff(c *db.Controller, facility constants.Facility) bool {
	return IsDivisionStaff(c) ||
		HasRole(c, constants.AirTrafficManager, facility) ||
		HasRole(c, constants.DeputyAirTrafficManager, facility) ||
		HasRole(c, constants.TrainingAdministrator, facility)
}

func IsATMOrDATM(c *db.Controller, facility constants.Facility) bool {
	return IsDivisionStaff(c) ||
		HasRole(c, constants.AirTrafficManager, facility) ||
		HasRole(c, constants.DeputyAirTrafficManager, facility)
}

func IsDivisionStaff(c *db.Controller) bool {
	return HasRole(c, constants.DivisionStaff, constants.Headquarters)
}

func IsInstructor(c *db.Controller, facility constants.Facility) bool {
	return IsDivisionStaff(c) ||
		HasRole(c, constants.Instructor, facility)
}

func IsTrainingStaff(c *db.Controller, facility constants.Facility) bool {
	return IsSeniorStaff(c, facility) ||
		HasRole(c, constants.Instructor, facility) ||
		HasRole(c, constants.Mentor, facility)
}
