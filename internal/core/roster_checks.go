package core

import (
	"github.com/VATUSA/api-v3/pkg/constants"
	db "github.com/VATUSA/api-v3/pkg/database"
)

func IsHomeController(c *db.Controller, facility constants.Facility) bool {
	return facility == c.Facility
}

func IsVisiting(c *db.Controller, fac constants.Facility) bool {
	for _, v := range c.Visits {
		if v.Facility == fac {
			return true
		}
	}
	return false
}

func IsOnFacilityRoster(c *db.Controller, facility constants.Facility) bool {
	return IsHomeController(c, facility) || IsVisiting(c, facility)
}
