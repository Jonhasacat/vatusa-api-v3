package database

import "gorm.io/gorm"

type EvaluationFormCategory struct {
	gorm.Model
	EvaluationFormID uint
	EvaluationForm   *EvaluationForm
	ParentCategoryID *uint
	ParentCategory   *EvaluationFormCategory
	Name             string
}
