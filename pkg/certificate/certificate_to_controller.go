package certificate

import (
	"fmt"
	"github.com/VATUSA/api-v3/pkg/action_log"
	"github.com/VATUSA/api-v3/pkg/database"
	"github.com/VATUSA/api-v3/pkg/facility"
	"github.com/VATUSA/api-v3/pkg/hold"
	"github.com/VATUSA/api-v3/pkg/rating"
	"github.com/VATUSA/api-v3/pkg/roster"
	"time"
)

func ControllerCertificateUpdated(c *database.Controller, certificate *database.Certificate) error {
	if c.ATCRating != certificate.Rating {
		if rating.Rating(certificate.Rating) >= rating.Observer &&
			rating.Rating(certificate.Rating) <= rating.I3 {
			oldRating := rating.Rating(c.ATCRating)
			newRating := rating.Rating(certificate.Rating)
			err := rating.ChangeRating(c, newRating, nil)
			if err != nil {
				return err
			}
			err = action_log.LogAction(c, fmt.Sprintf("Rating changed externally: %s -> %s",
				rating.ShortMap[oldRating], rating.ShortMap[newRating]), nil)
			if err != nil {
				return err
			}
		}
	}
	if rating.Rating(certificate.Rating) == rating.Inactive {
		if c.Facility != facility.Inactive {
			err := roster.ForceTransfer(c, facility.Inactive, "Inactive")
			if err != nil {
				return err
			}
		}
		c.IsInDivision = false
		c.IsActive = false
	} else if rating.Rating(certificate.Rating) == rating.Suspended {
		if facility.IsRosterFacility(c.Facility) {
			err := roster.RemoveFromFacility(c, nil, "Suspended")
			if err != nil {
				return err
			}
		} else {
			err := roster.RemoveAllVisits(c, nil, "Suspended")
			if err != nil {
				return err
			}
		}
		c.IsInDivision = false
		c.IsActive = false
	} else if *certificate.Division == "USA" && !c.IsInDivision {
		// Joined Division
		if c.Facility == facility.Inactive || c.Facility == facility.NonMember {
			err := roster.ForceTransfer(c, facility.Academy, "Joined division")
			if err != nil {
				return err
			}
		}
		err := action_log.LogAction(c, "Joined division", nil)
		if err != nil {
			return err
		}
		c.IsInDivision = true
		c.IsActive = true
		// TODO: Send Welcome to Division Notification
	} else if *certificate.Division != "USA" && c.IsInDivision {
		// Left Division
		if facility.IsRosterFacility(c.Facility) {
			err := roster.RemoveFromFacility(c, nil, "Left Division")
			if err != nil {
				return err
			}
		}
		err := action_log.LogAction(c, "Left division", nil)
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
		c.Facility = facility.Academy
		c.IsInDivision = true
		if c.ATCRating == rating.Observer {
			err := hold.AddHold(c, hold.Academy, "New Member", nil)
			if err != nil {
				return err
			}
		} else {
			err := hold.AddHold(c, hold.RCEExam, "New Member", nil)
			if err != nil {
				return err
			}
		}
	} else {
		c.Facility = facility.NonMember
		c.IsInDivision = false
		err := hold.AddHold(c, hold.RCEExam, "New Visitor", nil)
		if err != nil {
			return err
		}
	}
	if c.ATCRating > rating.I3 {
		c.ATCRating = rating.C1
	} else if c.ATCRating < rating.Observer {
		c.ATCRating = rating.Observer
	}
	err := c.Save()
	if err != nil {
		return err
	}
	// TODO
	return nil
}
