package core

import (
	"errors"
	"fmt"
	"github.com/VATUSA/api-v3/pkg/constants"
	"github.com/VATUSA/api-v3/pkg/database"
)

func ChangeRating(c *database.Controller, rating constants.Rating, requester *database.Controller) error {
	if rating == c.ATCRating {
		return errors.New(fmt.Sprintf(
			"Controller %d already has rating %d", c.CertificateId, c.ATCRating))
	}
	if rating < constants.Observer || rating > constants.I3 {
		return errors.New(fmt.Sprintf("Rating %d cannot be issued through this API!", rating))
	}
	if c.Certificate.Rating != rating && c.Certificate.Rating >= constants.Observer && c.Certificate.Rating < constants.Supervisor {
		// TODO: Change rating via VATSIM API
	}
	c.ATCRating = rating
	var requesterID uint64
	if requester != nil {
		requesterID = requester.Id
	} else {
		requesterID = 0
	}

	ratingChange := database.RatingChange{
		ControllerID: c.CertificateId,
		FromRating:   c.ATCRating,
		ToRating:     rating,
		AdminID:      requesterID,
	}
	database.DB.Create(&ratingChange)
	c.Save()
	return nil
}

func Promote(c *database.Controller, rating constants.Rating, requester *database.Controller) error {
	if constants.Rating(c.ATCRating) >= constants.C1 {
		return errors.New(fmt.Sprintf("Controller %d is already C1 or higher!", c.CertificateId))
	}
	if !IsPromotionEligible(c) {
		return errors.New(fmt.Sprintf("Controller %d is not eligible for promotion", c.CertificateId))
	}
	err := ChangeRating(c, rating, requester)
	if err != nil {
		return err
	}
	return nil
}
