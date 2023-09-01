package database

import (
	"gorm.io/gorm"
)

type ControllerRole struct {
	gorm.Model
	ControllerID uint64
	Controller   Controller
	Role         string
	Facility     string
}
