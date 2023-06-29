package database

import (
	"gorm.io/gorm"
)

type APITokenLog struct {
	gorm.Model
	APITokenID   uint
	APIToken     *APIToken
	ControllerID *uint64
	Controller   *Controller
	Message      string
}
