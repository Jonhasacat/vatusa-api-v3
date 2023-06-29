package operations

import (
	"errors"
	"fmt"
	"vatusa-api-v3/constants"
	db "vatusa-api-v3/database"
)

func ChangeRating(controller *db.Controller, rating constants.Rating, requester *db.Controller) error {
	if rating == controller.ATCRating {
		return errors.New(fmt.Sprintf(
			"Controller %d already has rating %d", controller.CertificateId, controller.ATCRating))
	}
	if rating < constants.RatingObserver || rating > constants.RatingI3 {
		return errors.New(fmt.Sprintf("Rating %d cannot be issued through this API!", rating))
	}
	if controller.Certificate.Rating != rating &&
		controller.Certificate.Rating >= constants.RatingObserver &&
		controller.Certificate.Rating < constants.RatingSupervisor {
		// TODO: Change rating via VATSIM API
	}
	controller.ATCRating = rating

	ratingChange := db.RatingChange{
		ControllerID: controller.CertificateId,
		FromRating:   controller.ATCRating,
		ToRating:     rating,
		AdminID:      requester.Id,
	}
	db.DB.Create(&ratingChange)
	controller.Save()
	return nil
}

func Promote(controller *db.Controller, rating constants.Rating, requester *db.Controller) error {
	if controller.ATCRating >= constants.RatingC1 {
		return errors.New(fmt.Sprintf("Controller %d is already C1 or higher!", controller.CertificateId))
	}
	if !controller.IsPromotionEligible() {
		return errors.New(fmt.Sprintf("Controller %d is not eligible for promotion", controller.CertificateId))
	}
	err := ChangeRating(controller, rating, requester)
	if err != nil {
		return err
	}
	return nil
}
