package roster

import (
	db "github.com/VATUSA/api-v3/pkg/database"
	"github.com/VATUSA/api-v3/pkg/facility"
)

func IsHomeController(c *db.Controller, facility facility.Facility) bool {
	return facility == c.Facility
}

func IsVisiting(c *db.Controller, fac facility.Facility) bool {
	for _, v := range c.Visits {
		if v.Facility == fac {
			return true
		}
	}
	return false
}

func IsOnFacilityRoster(c *db.Controller, facility facility.Facility) bool {
	return IsHomeController(c, facility) || IsVisiting(c, facility)
}
