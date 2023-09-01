package database

import (
	"github.com/VATUSA/api-v3/pkg/constants"
	"gorm.io/gorm"
)

type APIUser struct {
	gorm.Model
	Name     string
	Facility constants.Facility
}

func CreateAPIUser(name string, facility constants.Facility) (*APIUser, error) {
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
