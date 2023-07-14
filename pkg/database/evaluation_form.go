package database

import (
	"gorm.io/gorm"
)

type EvaluationForm struct {
	gorm.Model
	Name                  string
	StudentRequiredRating int
	StudentGrantRating    *int
	IsRatingEvaluation    bool
}
