package database

import (
	"gorm.io/gorm"
	"vatusa-api-v3/constants"
)

type EvaluationForm struct {
	gorm.Model
	Name                  string
	StudentRequiredRating constants.Rating
	StudentGrantRating    *constants.Rating
	IsRatingEvaluation    bool
}
