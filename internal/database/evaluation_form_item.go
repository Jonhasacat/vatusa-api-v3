package database

import "gorm.io/gorm"

type EvaluationFormItem struct {
	gorm.Model
	EvaluationFormID         uint
	EvaluationForm           *EvaluationForm
	EvaluationFormCategoryID uint
	EvaluationFormCategory   *EvaluationFormCategory
	Prompt                   string
	InstructorNote           string
	IsRequired               bool
}
