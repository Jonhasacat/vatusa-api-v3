package database

import "gorm.io/gorm"

type APIUser struct {
	gorm.Model
	Name     string
	Facility string
}

func CreateAPIUser(name string, facility string) (*APIUser, error) {
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
