package core

import (
	"errors"
	"fmt"
	database2 "github.com/VATUSA/api-v3/internal/database"
	"github.com/VATUSA/api-v3/pkg/constants"
)

func ChangeRating(c *database2.Controller, rating constants.Rating, requester *database2.Controller) error {
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

	ratingChange := database2.RatingChange{
		ControllerID: c.CertificateId,
		FromRating:   c.ATCRating,
		ToRating:     rating,
		AdminID:      requesterID,
	}
	database2.DB.Create(&ratingChange)
	c.Save()
	return nil
}

func Promote(c *database2.Controller, rating constants.Rating, requester *database2.Controller) error {
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
