package database

import (
	"github.com/VATUSA/api-v3/pkg/facility"
	"gorm.io/gorm"
)

type Transfer struct {
	gorm.Model
	ControllerID uint64
	Controller   *Controller
	FromFacility facility.Facility `gorm:"size:4"`
	ToFacility   facility.Facility `gorm:"size:4"`
	Reason       string            `gorm:"size:255"`
}

func (t *Transfer) Save() {
	DB.Save(t)
}
