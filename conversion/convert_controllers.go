package conversion

import (
	"time"
	"vatusa-api-v3/conversion/legacydb"
	"vatusa-api-v3/database"
)

func LoadLegacyControllers() ([]legacydb.Controller, error) {
	var controllers []legacydb.Controller
	result := legacydb.DB.Model(legacydb.Controller{}).Find(&controllers)
	if result.Error != nil {
		return nil, result.Error
	}
	return controllers, nil
}

func ProcessLegacyController(legacy legacydb.Controller) error {
	certificate, _ := database.FetchCertificateByID(legacy.CID)
	if certificate == nil {
		certificate = &database.Certificate{
			ID:                     legacy.CID,
			FirstName:              legacy.FName,
			LastName:               legacy.LName,
			Email:                  legacy.Email,
			Rating:                 legacy.Rating,
			PilotRating:            0,
			MilitaryRating:         0,
			SuspendDate:            nil,
			RegistrationDate:       nil,
			Region:                 nil,
			Division:               nil,
			SubDivision:            nil,
			LastRatingChange:       nil,
			CertificateUpdateStamp: time.Now(),
		}
		certificate.Save()
	}
	controller, _ := database.FetchControllerByCID(legacy.CID)
	if controller == nil {
		controller := database.Controller{
			Id:                        legacy.CID,
			CertificateId:             legacy.CID,
			Certificate:               certificate,
			Facility:                  legacy.Facility,
			FacilityJoin:              legacy.FacilityJoin,
			ATCRating:                 legacy.Rating,
			LastPromotion:             nil,
			IsInDivision:              legacy.FlagHomeController,
			IsApprovedExternalVisitor: false,
			IsActive:                  false,
			DiscordId:                 legacy.DiscordId,
		}
		controller.Save()
	} else {

	}
	return nil
}

func ConvertControllers() error {
	controllers, err := LoadLegacyControllers()
	if err != nil {
		return err
	}
	for _, controller := range controllers {
		err = ProcessLegacyController(controller)
		if err != nil {
			print(err.Error())
		}
	}
	return nil
}
