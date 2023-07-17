package role

import (
	db "github.com/VATUSA/api-v3/pkg/database"
	"github.com/VATUSA/api-v3/pkg/facility"
)

func HasRole(c *db.Controller, role Role, facility facility.Facility) bool {
	for _, r := range c.Roles {
		if r.Role == role && r.Facility == facility {
			return true
		}
	}
	return false
}

func IsStaff(c *db.Controller, facility facility.Facility) bool {
	return IsSeniorStaff(c, facility) ||
		HasRole(c, WebMaster, facility) ||
		HasRole(c, FacilityEngineer, facility) ||
		HasRole(c, EventCoordinator, facility) ||
		IsTrainingStaff(c, facility)
}

func IsSeniorStaff(c *db.Controller, facility facility.Facility) bool {
	return IsDivisionStaff(c) ||
		HasRole(c, AirTrafficManager, facility) ||
		HasRole(c, DeputyAirTrafficManager, facility) ||
		HasRole(c, TrainingAdministrator, facility)
}

func IsATMOrDATM(c *db.Controller, facility facility.Facility) bool {
	return IsDivisionStaff(c) ||
		HasRole(c, AirTrafficManager, facility) ||
		HasRole(c, DeputyAirTrafficManager, facility)
}

func IsDivisionStaff(c *db.Controller) bool {
	return HasRole(c, DivisionStaff, facility.Headquarters)
}

func IsInstructor(c *db.Controller, facility facility.Facility) bool {
	return IsDivisionStaff(c) ||
		HasRole(c, Instructor, facility)
}

func IsTrainingStaff(c *db.Controller, facility facility.Facility) bool {
	return IsSeniorStaff(c, facility) ||
		HasRole(c, Instructor, facility) ||
		HasRole(c, Mentor, facility)
}
