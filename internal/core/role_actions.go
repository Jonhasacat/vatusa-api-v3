package core

import (
	"errors"
	"fmt"
	"github.com/VATUSA/api-v3/internal/database"
	"github.com/VATUSA/api-v3/pkg/constants"
)

func AddRole(c *database.Controller, r constants.Role, f constants.Facility, requester *database.Controller) error {
	if HasRole(c, r, f) {
		return errors.New(fmt.Sprintf("controller %d already has role %s for facility %s", c.Id, r, f))
	}
	record := &database.ControllerRole{
		ControllerID: c.Id,
		Controller:   *c,
		Role:         r,
		Facility:     r,
	}
	result := database.DB.Save(record)
	if result.Error != nil {
		return result.Error
	}
	err := LogAction(c, fmt.Sprintf("Added role %s for facility %s", r, f), requester)
	if err != nil {
		return err
	}
	return nil
}

func RemoveRole(c *database.Controller, r constants.Role, f constants.Facility, requester *database.Controller) error {
	for _, role := range c.Roles {
		if role.Role == r && role.Facility == f {
			database.DB.Delete(&role)
			err := LogAction(c, fmt.Sprintf("Removed role %s for facility %s", r, f), requester)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Controller %d does not have role %s for facility %s", c.Id, r, f))
}

func RemoveFacilityRoles(c *database.Controller, f constants.Facility) error {
	for _, r := range c.Roles {
		if constants.Facility(r.Facility) == f {
			result := database.DB.Delete(r)
			if result.Error != nil {
				return result.Error
			}
		}
	}
	return nil
}
