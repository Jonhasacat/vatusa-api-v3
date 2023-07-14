package controller

import (
	db "github.com/VATUSA/api-v3/pkg/database"
	"github.com/VATUSA/api-v3/pkg/facility"
	"github.com/VATUSA/api-v3/pkg/role"
)

func HasRole(c *db.Controller, role role.Role, facility facility.Facility) bool {
	for _, r := range c.Roles {
		if r.Role == string(role) && r.Facility == string(facility) {
			return true
		}
	}
	return false
}

func IsStaff(c *db.Controller, facility facility.Facility) bool {
	return IsSeniorStaff(c, facility) ||
		HasRole(c, role.RoleWebMaster, facility) ||
		HasRole(c, role.RoleFacilityEngineer, facility) ||
		HasRole(c, role.RoleEventCoordinator, facility) ||
		IsTrainingStaff(c, facility)
}

func IsSeniorStaff(c *db.Controller, facility facility.Facility) bool {
	return IsDivisionStaff(c) ||
		HasRole(c, role.RoleAirTrafficManager, facility) ||
		HasRole(c, role.RoleDeputyAirTrafficManager, facility) ||
		HasRole(c, role.RoleTrainingAdministrator, facility)
}

func IsATMOrDATM(c *db.Controller, facility facility.Facility) bool {
	return IsDivisionStaff(c) ||
		HasRole(c, role.RoleAirTrafficManager, facility) ||
		HasRole(c, role.RoleDeputyAirTrafficManager, facility)
}

func IsDivisionStaff(c *db.Controller) bool {
	return HasRole(c, role.RoleDivisionStaff, facility.Headquarters)
}

func IsInstructor(c *db.Controller, facility facility.Facility) bool {
	return IsDivisionStaff(c) ||
		HasRole(c, role.RoleInstructor, facility)
}

func IsTrainingStaff(c *db.Controller, facility facility.Facility) bool {
	return IsSeniorStaff(c, facility) ||
		HasRole(c, role.RoleInstructor, facility) ||
		HasRole(c, role.RoleMentor, facility)
}
