package core

import (
	"fmt"
	"github.com/VATUSA/api-v3/pkg/constants"
	"github.com/VATUSA/api-v3/pkg/database"
	"time"
)

func ControllerCertificateUpdated(c *database.Controller, certificate *database.Certificate) error {
	if c.ATCRating != certificate.Rating {
		if constants.Rating(certificate.Rating) >= constants.Observer &&
			constants.Rating(certificate.Rating) <= constants.I3 {
			oldRating := constants.Rating(c.ATCRating)
			newRating := constants.Rating(certificate.Rating)
			err := ChangeRating(c, newRating, nil)
			if err != nil {
				return err
			}
			err = LogAction(c, fmt.Sprintf("Rating changed externally: %s -> %s",
				constants.ShortMap[oldRating], constants.ShortMap[newRating]), nil)
			if err != nil {
				return err
			}
		}
	}
	if constants.Rating(certificate.Rating) == constants.Inactive {
		if c.Facility != constants.InactiveFacility {
			err := ForceTransfer(c, constants.InactiveFacility, "Inactive")
			if err != nil {
				return err
			}
		}
		c.IsInDivision = false
		c.IsActive = false
	} else if constants.Rating(certificate.Rating) == constants.Suspended {
		if constants.IsRosterFacility(c.Facility) {
			err := RemoveFromFacility(c, nil, "Suspended")
			if err != nil {
				return err
			}
		} else {
			err := RemoveAllVisits(c, nil, "Suspended")
			if err != nil {
				return err
			}
		}
		c.IsInDivision = false
		c.IsActive = false
	} else if *certificate.Division == "USA" && !c.IsInDivision {
		// Joined Division
		if c.Facility == constants.InactiveFacility || c.Facility == constants.NonMember {
			err := ForceTransfer(c, constants.Academy, "Joined division")
			if err != nil {
				return err
			}
		}
		err := LogAction(c, "Joined division", nil)
		if err != nil {
			return err
		}
		c.IsInDivision = true
		c.IsActive = true
		// TODO: Send Welcome to Division Notification
	} else if *certificate.Division != "USA" && c.IsInDivision {
		// Left Division
		if constants.IsRosterFacility(c.Facility) {
			err := RemoveFromFacility(c, nil, "Left Division")
			if err != nil {
				return err
			}
		}
		err := LogAction(c, "Left division", nil)
		if err != nil {
			return err
		}
		c.IsInDivision = false
		c.IsActive = true
	}
	c.Save()
	// TODO: Send Exit Survey Notification?
	return nil
}

func NewController(certificate *database.Certificate) error {
	now := time.Now()
	c := &database.Controller{
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
		c.Facility = constants.Academy
		c.IsInDivision = true
		if c.ATCRating == constants.Observer {
			err := AddHold(c, constants.Academy, "New Member", nil)
			if err != nil {
				return err
			}
		} else {
			err := AddHold(c, constants.RCEExam, "New Member", nil)
			if err != nil {
				return err
			}
		}
	} else {
		c.Facility = constants.NonMember
		c.IsInDivision = false
		err := AddHold(c, constants.RCEExam, "New Visitor", nil)
		if err != nil {
			return err
		}
	}
	if c.ATCRating > constants.I3 {
		c.ATCRating = constants.C1
	} else if c.ATCRating < constants.Observer {
		c.ATCRating = constants.Observer
	}
	err := c.Save()
	if err != nil {
		return err
	}
	// TODO
	return nil
}
