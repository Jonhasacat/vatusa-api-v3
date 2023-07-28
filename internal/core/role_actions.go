package core

import (
	"errors"
	"fmt"
	"github.com/VATUSA/api-v3/pkg/constants"
	db "github.com/VATUSA/api-v3/pkg/database"
)

func AddRole(c *db.Controller, r constants.Role, f constants.Facility, requester *db.Controller) error {
	if HasRole(c, r, f) {
		return errors.New(fmt.Sprintf("controller %d already has role %s for facility %s", c.Id, r, f))
	}
	record := &db.ControllerRole{
		ControllerID: c.Id,
		Controller:   *c,
		Role:         r,
		Facility:     r,
	}
	result := db.DB.Save(record)
	if result.Error != nil {
		return result.Error
	}
	err := LogAction(c, fmt.Sprintf("Added role %s for facility %s", r, f), requester)
	if err != nil {
		return err
	}
	return nil
}

func RemoveRole(c *db.Controller, r constants.Role, f constants.Facility, requester *db.Controller) error {
	for _, role := range c.Roles {
		if role.Role == r && role.Facility == f {
			db.DB.Delete(&role)
			err := LogAction(c, fmt.Sprintf("Removed role %s for facility %s", r, f), requester)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Controller %d does not have role %s for facility %s", c.Id, r, f))
}

func RemoveFacilityRoles(c *db.Controller, f constants.Facility) error {
	for _, r := range c.Roles {
		if constants.Facility(r.Facility) == f {
			result := db.DB.Delete(r)
			if result.Error != nil {
				return result.Error
			}
		}
	}
	return nil
}
