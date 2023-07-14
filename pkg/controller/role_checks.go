package controller

import (
	db "github.com/VATUSA/api-v3/pkg/database"
	"github.com/VATUSA/api-v3/pkg/facility"
	"github.com/VATUSA/api-v3/pkg/role"
)

func HasRole(c *db.Controller, role role.Role, facility facility.Facility) bool {
	for _, r := range c.Roles {
		if r.Role == role && r.Facility == facility {
			return true
		}
	}
	return false
}

func IsStaff(c *db.Controller, facility facility.Facility) bool {
	return IsSeniorStaff(c, facility) ||
		HasRole(c, role.WebMaster, facility) ||
		HasRole(c, role.FacilityEngineer, facility) ||
		HasRole(c, role.EventCoordinator, facility) ||
		IsTrainingStaff(c, facility)
}

func IsSeniorStaff(c *db.Controller, facility facility.Facility) bool {
	return IsDivisionStaff(c) ||
		HasRole(c, role.AirTrafficManager, facility) ||
		HasRole(c, role.DeputyAirTrafficManager, facility) ||
		HasRole(c, role.TrainingAdministrator, facility)
}

func IsATMOrDATM(c *db.Controller, facility facility.Facility) bool {
	return IsDivisionStaff(c) ||
		HasRole(c, role.AirTrafficManager, facility) ||
		HasRole(c, role.DeputyAirTrafficManager, facility)
}

func IsDivisionStaff(c *db.Controller) bool {
	return HasRole(c, role.DivisionStaff, facility.Headquarters)
}

func IsInstructor(c *db.Controller, facility facility.Facility) bool {
	return IsDivisionStaff(c) ||
		HasRole(c, role.Instructor, facility)
}

func IsTrainingStaff(c *db.Controller, facility facility.Facility) bool {
	return IsSeniorStaff(c, facility) ||
		HasRole(c, role.Instructor, facility) ||
		HasRole(c, role.Mentor, facility)
}
