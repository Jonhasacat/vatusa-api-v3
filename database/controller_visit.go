package database

import (
	"gorm.io/gorm"
	"vatusa-api-v3/constants"
)

type ControllerVisit struct {
	gorm.Model
	ControllerID uint64
	Controller   *Controller
	Facility     constants.Facility
}

func (c *ControllerVisit) Save() {
	DB.Save(c)
	c.Controller.HookControllerUpdate()
}
