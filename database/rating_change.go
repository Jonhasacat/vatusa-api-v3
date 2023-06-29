package database

import (
	"gorm.io/gorm"
	"vatusa-api-v3/constants"
)

type RatingChange struct {
	gorm.Model
	ControllerID uint64
	FromRating   constants.Rating
	ToRating     constants.Rating
	AdminID      uint64
}
