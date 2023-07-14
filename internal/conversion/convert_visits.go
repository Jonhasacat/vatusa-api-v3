package conversion

import (
	"github.com/VATUSA/api-v3/internal/conversion/legacydb"
	"github.com/VATUSA/api-v3/pkg/controller"
	db "github.com/VATUSA/api-v3/pkg/database"
	"github.com/VATUSA/api-v3/pkg/facility"
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
	c, err := db.FetchControllerByCID(legacy.CID)
	if err != nil {
		return err
	}
	if !controller.IsVisiting(c, facility.Facility(legacy.Facility)) {
		visit := db.ControllerVisit{
			Model: gorm.Model{
				CreatedAt: legacy.CreatedAt,
				UpdatedAt: legacy.UpdatedAt,
			},
			ControllerID: c.Id,
			Controller:   c,
			Facility:     facility.Facility(legacy.Facility),
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
