package database

import (
	"gorm.io/gorm"
	"vatusa-api-v3/constants"
)

type Transfer struct {
	gorm.Model
	ControllerID uint64
	Controller   *Controller
	FromFacility constants.Facility `gorm:"size:4"`
	ToFacility   constants.Facility `gorm:"size:4"`
	Reason       string             `gorm:"size:255"`
}

func (t *Transfer) Save() {
	DB.Save(t)
}
