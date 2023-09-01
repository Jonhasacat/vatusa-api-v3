package database

import "gorm.io/gorm"

type WorkspaceDomain struct {
	gorm.Model
	Domain    string
	Facility  string
	IsEnabled bool
}
