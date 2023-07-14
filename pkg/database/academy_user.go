package database

import (
	"gorm.io/gorm"
	"time"
)

type AcademyUser struct {
	gorm.Model
	ControllerID uint64
	Controller   *Controller
	MoodleUserID uint64
	LastSync     time.Time
}
