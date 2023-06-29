package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type APIToken struct {
	gorm.Model
	Token     string
	APIUserID uint
	APIUser   *APIUser
	IsEnabled bool
}

func tokenQuery() *gorm.DB {
	return DB.Model(&APIToken{}).Joins("APIUser")
}

func FetchAPITokenByToken(token string) (*APIToken, error) {
	var record APIToken
	result := tokenQuery().Where("Token = ?", token).First(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return &record, nil
}

func GenerateAPIToken(user *APIUser, controller *Controller) (*APIToken, error) {
	token := APIToken{
		Token:     uuid.New().String(),
		APIUser:   user,
		IsEnabled: true,
	}
	result := DB.Save(&token)
	if result.Error != nil {
		return nil, result.Error
	}
	err := token.Log("Created Token", controller)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (t *APIToken) Log(message string, controller *Controller) error {
	log := APITokenLog{
		APIToken:   t,
		Controller: controller,
		Message:    message,
	}
	result := DB.Save(&log)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
