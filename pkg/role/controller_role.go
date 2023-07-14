package role

import (
	"github.com/VATUSA/api-v3/pkg/database"
	"github.com/VATUSA/api-v3/pkg/facility"
)

func AddRole(c *database.Controller, r Role, f facility.Facility) error {
	// TODO
	return nil
}

func RemoveRole(c *database.Controller, r Role, f facility.Facility) error {
	// TODO
	return nil
}

func RemoveFacilityRoles(c *database.Controller, f facility.Facility) error {
	for _, r := range c.Roles {
		if facility.Facility(r.Facility) == f {
			result := database.DB.Delete(r)
			if result.Error != nil {
				return result.Error
			}
		}
	}
	return nil
}
