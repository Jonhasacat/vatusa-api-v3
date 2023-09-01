package core

import (
	"errors"
	"github.com/VATUSA/api-v3/internal/database"
	"github.com/VATUSA/api-v3/pkg/vatsim"
	"gorm.io/gorm"
	"time"
)

func CreateCertificate(member *vatsim.Member) error {
	cert := &database.Certificate{
		ID:                     member.ID,
		FirstName:              *member.NameFirst,
		LastName:               *member.NameLast,
		Email:                  *member.Email,
		Rating:                 member.Rating,
		PilotRating:            member.PilotRating,
		MilitaryRating:         member.MilitaryRating,
		SuspendDate:            member.SuspendTime(),
		RegistrationDate:       member.RegistrationTime(),
		Region:                 &member.Region,
		Division:               &member.Division,
		SubDivision:            &member.SubDivision,
		LastRatingChange:       member.LastRatingChangeTime(),
		CertificateUpdateStamp: time.Now(),
	}
	result := database.DB.Create(cert)
	if result.Error != nil {
		return result.Error
	}
	controller, err := database.FetchControllerByCID(cert.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if controller != nil {
		err = ControllerCertificateUpdated(controller, cert)
		if err != nil {
			return err
		}
	} else if member.Region == "AMAS" && member.Division == "USA" {
		_, err = NewController(cert)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateCertificate(cert *database.Certificate, member *vatsim.Member) error {
	if member.NameFirst != nil {
		cert.FirstName = *member.NameFirst
	}
	if member.NameLast != nil {
		cert.LastName = *member.NameLast
	}
	if member.Email != nil {
		cert.Email = *member.Email
	}
	cert.Rating = member.Rating
	cert.PilotRating = member.PilotRating
	cert.MilitaryRating = member.MilitaryRating
	cert.SuspendDate = member.SuspendTime()
	cert.RegistrationDate = member.RegistrationTime()
	cert.Region = &member.Region
	cert.Division = &member.Division
	cert.SubDivision = &member.SubDivision
	cert.LastRatingChange = member.LastRatingChangeTime()
	cert.CertificateUpdateStamp = time.Now()
	err := cert.Save()
	if err != nil {
		return err
	}
	controller, err := database.FetchControllerByCID(cert.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if controller != nil {
		err = ControllerCertificateUpdated(controller, cert)
		if err != nil {
			return err
		}
	} else if member.Region == "AMAS" && member.Division == "USA" {
		_, err = NewController(cert)
		if err != nil {
			return err
		}
	}
	return nil
}
