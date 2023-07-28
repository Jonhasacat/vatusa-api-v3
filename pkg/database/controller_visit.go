package database

import (
	"github.com/VATUSA/api-v3/pkg/constants"
	"gorm.io/gorm"
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
