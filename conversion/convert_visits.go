package conversion

import (
	"gorm.io/gorm"
	"vatusa-api-v3/conversion/legacydb"
	"vatusa-api-v3/database"
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
	controller, err := database.FetchControllerByCID(legacy.CID)
	if err != nil {
		return err
	}
	if !controller.IsVisiting(legacy.Facility) {
		visit := database.ControllerVisit{
			Model: gorm.Model{
				CreatedAt: legacy.CreatedAt,
				UpdatedAt: legacy.UpdatedAt,
			},
			ControllerID: controller.Id,
			Controller:   controller,
			Facility:     legacy.Facility,
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
