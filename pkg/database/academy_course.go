package database

import "gorm.io/gorm"

type AcademyCourse struct {
	gorm.Model
	DisplayName           string
	CourseType            string
	MoodleCourseID        uint64
	MoodleEnrolID         *uint64
	Facility              string
	StudentRequiredRating uint64
	AllowManualEnroll     bool
	IsExamOnly            bool
}
