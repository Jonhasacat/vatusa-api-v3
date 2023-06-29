package operations

import (
	"fmt"
	"time"
	"vatusa-api-v3/constants"
	db "vatusa-api-v3/database"
)

func NewController(certificate *db.Certificate) error {
	now := time.Now()
	controller := &db.Controller{
		Id:                        certificate.ID,
		CertificateId:             certificate.ID,
		Certificate:               certificate,
		FacilityJoin:              &now,
		ATCRating:                 certificate.Rating,
		LastPromotion:             nil,
		IsApprovedExternalVisitor: false,
		IsActive:                  true,
		DiscordId:                 nil,
	}
	if *certificate.Division == "USA" {
		controller.Facility = constants.Academy
		controller.IsInDivision = true
		if controller.ATCRating == constants.RatingObserver {
			err := AddHold(controller, constants.HoldAcademy, "New Member", nil)
			if err != nil {
				return err
			}
		} else {
			err := AddHold(controller, constants.HoldRCEExam, "New Member", nil)
			if err != nil {
				return err
			}
		}
	} else {
		controller.Facility = constants.NonMember
		controller.IsInDivision = false
		err := AddHold(controller, constants.HoldRCEExam, "New Visitor", nil)
		if err != nil {
			return err
		}
	}
	if controller.ATCRating > constants.RatingI3 {
		controller.ATCRating = constants.RatingC1
	} else if controller.ATCRating < constants.RatingObserver {
		controller.ATCRating = constants.RatingObserver
	}
	result := db.DB.Create(controller)
	if result.Error != nil {
		return result.Error
	}
	// TODO
	return nil
}

func CertificateUpdated(controller *db.Controller, certificate *db.Certificate) error {
	if controller.ATCRating != certificate.Rating {
		if certificate.Rating >= constants.RatingObserver &&
			certificate.Rating <= constants.RatingI3 {
			oldRating := controller.ATCRating
			newRating := certificate.Rating
			err := ChangeRating(controller, certificate.Rating, nil)
			if err != nil {
				return err
			}
			err = LogAction(controller, fmt.Sprintf("Rating changed externally: %s -> %s",
				constants.RatingShortMap[oldRating], constants.RatingShortMap[newRating]), nil)
			if err != nil {
				return err
			}
		}
	}
	if certificate.Rating == constants.RatingInactive {
		if controller.Facility != constants.Inactive {
			err := ForceTransfer(controller, constants.Inactive, "Inactive")
			if err != nil {
				return err
			}
		}
		controller.IsInDivision = false
		controller.IsActive = false
	} else if certificate.Rating == constants.RatingSuspended {
		if constants.IsRosterFacility(controller.Facility) {
			err := RemoveFromFacility(controller, nil, "Suspended")
			if err != nil {
				return err
			}
		} else {
			err := RemoveAllVisits(controller, nil, "Suspended")
			if err != nil {
				return err
			}
		}
		controller.IsInDivision = false
		controller.IsActive = false
	} else if *certificate.Division == "USA" && !controller.IsInDivision {
		// Joined Division
		if controller.Facility == constants.Inactive || controller.Facility == constants.NonMember {
			err := ForceTransfer(controller, constants.Academy, "Joined division")
			if err != nil {
				return err
			}
		}
		err := LogAction(controller, "Joined division", nil)
		if err != nil {
			return err
		}
		controller.IsInDivision = true
		controller.IsActive = true
		// TODO: Send Welcome to Division Notification
	} else if *certificate.Division != "USA" && controller.IsInDivision {
		// Left Division
		if constants.IsRosterFacility(controller.Facility) {
			err := RemoveFromFacility(controller, nil, "Left Division")
			if err != nil {
				return err
			}
		}
		err := LogAction(controller, "Left division", nil)
		if err != nil {
			return err
		}
		controller.IsInDivision = false
		controller.IsActive = true
	}
	controller.Save()
	// TODO: Send Exit Survey Notification?
	return nil
}
