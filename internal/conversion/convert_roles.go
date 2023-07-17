package conversion

import (
	"github.com/VATUSA/api-v3/internal/conversion/legacydb"
	db "github.com/VATUSA/api-v3/pkg/database"
	"github.com/VATUSA/api-v3/pkg/role"
)

func LoadLegacyRoles() ([]legacydb.Role, error) {
	var roles []legacydb.Role
	result := legacydb.DB.Model(legacydb.Role{}).Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}

func ProcessLegacyRole(legacyRole legacydb.Role) error {
	c, err := db.FetchControllerByCID(legacyRole.CID)
	if err != nil {
		return err
	}
	if !role.HasRole(c, legacyRole.Role, legacyRole.Facility) {
		record := &db.ControllerRole{
			ControllerID: c.Id,
			Controller:   *c,
			Role:         legacyRole.Role,
			Facility:     legacyRole.Facility,
		}
		result := db.DB.Save(record)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func ConvertRoles() error {
	roles, err := LoadLegacyRoles()
	if err != nil {
		return err
	}
	for _, r := range roles {
		err = ProcessLegacyRole(r)
		if err != nil {
			print(err.Error())
		}
	}
	return nil
}
