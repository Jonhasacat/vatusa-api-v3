package database

import "gorm.io/gorm"

type Evaluation struct {
	gorm.Model
	StudentId    uint64
	Student      *Controller `gorm:"foreignKey:StudentId"`
	InstructorId uint64
	Instructor   *Controller `gorm:"foreignKey:InstructorId"`
	Facility     string
	Position     string
	Notes        string
	IsPass       bool
}
