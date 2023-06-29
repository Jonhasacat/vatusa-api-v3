package operations

import (
	"vatusa-api-v3/constants"
	db "vatusa-api-v3/database"
)

func AddRole(controller *db.Controller, role constants.Role, facility constants.Facility) error {
	// TODO
	return nil
}

func RemoveRole(controller *db.Controller, role constants.Role, facility constants.Facility) error {
	// TODO
	return nil
}

func RemoveFacilityRoles(controller *db.Controller, facility constants.Facility) error {
	for _, r := range controller.Roles {
		if r.Facility == facility {
			result := db.DB.Delete(r)
			if result.Error != nil {
				return result.Error
			}
		}
	}
	return nil
}
