package facility

import (
	"fmt"
	"github.com/VATUSA/api-v3/pkg/action_log"
	controller2 "github.com/VATUSA/api-v3/pkg/controller"
	"github.com/VATUSA/api-v3/pkg/database"
	"github.com/VATUSA/api-v3/pkg/facility"
	"github.com/VATUSA/api-v3/pkg/hold"
	"github.com/VATUSA/api-v3/pkg/rating"
	"time"
)

func CertificateUpdated(controller *database.Controller, certificate *database.Certificate) error {
	if controller.ATCRating != certificate.Rating {
		if rating.Rating(certificate.Rating) >= rating.Observer &&
			rating.Rating(certificate.Rating) <= rating.I3 {
			oldRating := rating.Rating(controller.ATCRating)
			newRating := rating.Rating(certificate.Rating)
			err := rating.ChangeRating(controller, newRating, nil)
			if err != nil {
				return err
			}
			err = action_log.LogAction(controller, fmt.Sprintf("Rating changed externally: %s -> %s",
				rating.ShortMap[oldRating], rating.ShortMap[newRating]), nil)
			if err != nil {
				return err
			}
		}
	}
	if rating.Rating(certificate.Rating) == rating.Inactive {
		if controller.Facility != facility.Inactive {
			err := controller2.ForceTransfer(controller, facility.Inactive, "Inactive")
			if err != nil {
				return err
			}
		}
		controller.IsInDivision = false
		controller.IsActive = false
	} else if rating.Rating(certificate.Rating) == rating.Suspended {
		if facility.IsRosterFacility(controller.Facility) {
			err := controller2.RemoveFromFacility(controller, nil, "Suspended")
			if err != nil {
				return err
			}
		} else {
			err := controller2.RemoveAllVisits(controller, nil, "Suspended")
			if err != nil {
				return err
			}
		}
		controller.IsInDivision = false
		controller.IsActive = false
	} else if *certificate.Division == "USA" && !controller.IsInDivision {
		// Joined Division
		if controller.Facility == facility.Inactive || controller.Facility == facility.NonMember {
			err := controller2.ForceTransfer(controller, facility.Academy, "Joined division")
			if err != nil {
				return err
			}
		}
		err := action_log.LogAction(controller, "Joined division", nil)
		if err != nil {
			return err
		}
		controller.IsInDivision = true
		controller.IsActive = true
		// TODO: Send Welcome to Division Notification
	} else if *certificate.Division != "USA" && controller.IsInDivision {
		// Left Division
		if facility.IsRosterFacility(controller.Facility) {
			err := controller2.RemoveFromFacility(controller, nil, "Left Division")
			if err != nil {
				return err
			}
		}
		err := action_log.LogAction(controller, "Left division", nil)
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

func NewController(certificate *database.Certificate) error {
	now := time.Now()
	controller := &database.Controller{
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
		controller.Facility = facility.Academy
		controller.IsInDivision = true
		if rating.Rating(controller.ATCRating) == rating.Observer {
			err := hold.AddHold(controller, hold.Academy, "New Member", nil)
			if err != nil {
				return err
			}
		} else {
			err := hold.AddHold(controller, hold.RCEExam, "New Member", nil)
			if err != nil {
				return err
			}
		}
	} else {
		controller.Facility = facility.NonMember
		controller.IsInDivision = false
		err := hold.AddHold(controller, hold.RCEExam, "New Visitor", nil)
		if err != nil {
			return err
		}
	}
	if rating.Rating(controller.ATCRating) > rating.I3 {
		controller.ATCRating = int(rating.C1)
	} else if rating.Rating(controller.ATCRating) < rating.Observer {
		controller.ATCRating = int(rating.Observer)
	}
	result := database.DB.Create(controller)
	if result.Error != nil {
		return result.Error
	}
	// TODO
	return nil
}
