package database

import (
	"github.com/VATUSA/api-v3/pkg/facility"
	"gorm.io/gorm"
)

type APIUser struct {
	gorm.Model
	Name     string
	Facility facility.Facility
}

func CreateAPIUser(name string, facility facility.Facility) (*APIUser, error) {
	user := APIUser{
		Name:     name,
		Facility: facility,
	}
	result := DB.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
