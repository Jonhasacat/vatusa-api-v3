package database

import (
	"gorm.io/gorm"
)

type RatingChange struct {
	gorm.Model
	ControllerID uint64
	Controller   *Controller
	FromRating   int
	ToRating     int
	AdminID      uint64
}
