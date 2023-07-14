package conversion

import (
	legacydb2 "github.com/VATUSA/api-v3/internal/conversion/legacydb"
	database2 "github.com/VATUSA/api-v3/pkg/database"
	"time"
)

func LoadLegacyControllers() ([]legacydb2.Controller, error) {
	var controllers []legacydb2.Controller
	result := legacydb2.DB.Model(legacydb2.Controller{}).Find(&controllers)
	if result.Error != nil {
		return nil, result.Error
	}
	return controllers, nil
}

func ProcessLegacyController(legacy legacydb2.Controller) error {
	certificate, _ := database2.FetchCertificateByID(legacy.CID)
	if certificate == nil {
		certificate = &database2.Certificate{
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
	controller, _ := database2.FetchControllerByCID(legacy.CID)
	if controller == nil {
		controller := database2.Controller{
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
