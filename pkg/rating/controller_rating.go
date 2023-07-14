package rating

import (
	"errors"
	"fmt"
	"github.com/VATUSA/api-v3/pkg/controller"
	"github.com/VATUSA/api-v3/pkg/database"
)

func ChangeRating(c *database.Controller, rating Rating, requester *database.Controller) error {
	if rating == Rating(c.ATCRating) {
		return errors.New(fmt.Sprintf(
			"Controller %d already has rating %d", c.CertificateId, c.ATCRating))
	}
	if rating < Observer || rating > I3 {
		return errors.New(fmt.Sprintf("Rating %d cannot be issued through this API!", rating))
	}
	if Rating(c.Certificate.Rating) != rating &&
		Rating(c.Certificate.Rating) >= Observer &&
		Rating(c.Certificate.Rating) < Supervisor {
		// TODO: Change rating via VATSIM API
	}
	c.ATCRating = int(rating)

	ratingChange := database.RatingChange{
		ControllerID: c.CertificateId,
		FromRating:   c.ATCRating,
		ToRating:     int(rating),
		AdminID:      requester.Id,
	}
	database.DB.Create(&ratingChange)
	c.Save()
	return nil
}

func Promote(c *database.Controller, rating Rating, requester *database.Controller) error {
	if Rating(c.ATCRating) >= C1 {
		return errors.New(fmt.Sprintf("Controller %d is already C1 or higher!", c.CertificateId))
	}
	if !controller.IsPromotionEligible(c) {
		return errors.New(fmt.Sprintf("Controller %d is not eligible for promotion", c.CertificateId))
	}
	err := ChangeRating(c, rating, requester)
	if err != nil {
		return err
	}
	return nil
}
