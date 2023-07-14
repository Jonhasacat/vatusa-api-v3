package database

import (
	"gorm.io/gorm"
)

type Transfer struct {
	gorm.Model
	ControllerID uint64
	Controller   *Controller
	FromFacility string `gorm:"size:4"`
	ToFacility   string `gorm:"size:4"`
	Reason       string `gorm:"size:255"`
}

func (t *Transfer) Save() {
	DB.Save(t)
}
