package database

import (
	"gorm.io/gorm"
	"vatusa-api-v3/constants"
)

type ControllerRole struct {
	gorm.Model
	ControllerID uint64
	Controller   Controller
	Role         constants.Role
	Facility     constants.Facility
}
