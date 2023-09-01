package database

import "gorm.io/gorm"

type AcademyExam struct {
	gorm.Model
	AcademyCourseID  uint
	AcademyCourse    *AcademyCourse
	MoodleExamID     uint64
	RatingEquivalent *uint64
}
