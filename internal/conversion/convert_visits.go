package conversion

import (
	"github.com/VATUSA/api-v3/internal/conversion/legacydb"
	"github.com/VATUSA/api-v3/internal/core"
	"github.com/VATUSA/api-v3/internal/database"
	"github.com/VATUSA/api-v3/pkg/constants"
	"gorm.io/gorm"
)

func LoadLegacyVisits() ([]legacydb.Visit, error) {
	var visits []legacydb.Visit
	result := legacydb.DB.Model(legacydb.Visit{}).Find(&visits)
	if result.Error != nil {
		return nil, result.Error
	}
	return visits, nil
}

func ProcessLegacyVisit(legacy legacydb.Visit) error {
	c, err := database.FetchControllerByCID(legacy.CID)
	if err != nil {
		return err
	}
	if !core.IsVisiting(c, constants.Facility(legacy.Facility)) {
		visit := database.ControllerVisit{
			Model: gorm.Model{
				CreatedAt: legacy.CreatedAt,
				UpdatedAt: legacy.UpdatedAt,
			},
			ControllerID: c.Id,
			Controller:   c,
			Facility:     constants.Facility(legacy.Facility),
		}
		visit.Save()
	}
	return nil
}

func ConvertVisits() error {
	visits, err := LoadLegacyVisits()
	if err != nil {
		return err
	}
	for _, v := range visits {
		err = ProcessLegacyVisit(v)
		if err != nil {
			println(err.Error())
		}
	}
	return nil
}
