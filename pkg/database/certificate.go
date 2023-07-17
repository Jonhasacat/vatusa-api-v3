package database

import (
	"time"
)

type Certificate struct {
	ID                     uint64 `gorm:"primarykey;autoIncrement:false"`
	FirstName              string `gorm:"size:120"`
	LastName               string `gorm:"size:120"`
	Email                  string `gorm:"size:120"`
	Rating                 int
	PilotRating            int
	MilitaryRating         int
	SuspendDate            *time.Time
	RegistrationDate       *time.Time
	Region                 *string `gorm:"size:20"`
	Division               *string `gorm:"size:20"`
	SubDivision            *string `gorm:"size:20"`
	LastRatingChange       *time.Time
	CertificateUpdateStamp time.Time
}

func FetchCertificateByID(id uint64) (*Certificate, error) {
	var model *Certificate
	result := DB.Model(&Certificate{}).First(&model, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}

func (c *Certificate) Save() error {
	result := DB.Save(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
