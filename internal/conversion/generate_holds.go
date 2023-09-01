package conversion

import (
	"github.com/VATUSA/api-v3/internal/core"
	"github.com/VATUSA/api-v3/internal/database"
	"github.com/VATUSA/api-v3/pkg/constants"
	"time"
)

func GenerateHoldsForController(controller *database.Controller) error {
	// Academy / RCE
	if controller.Facility == constants.Academy && controller.ATCRating == constants.Observer {
		err := core.AddHold(controller, constants.Academy, "New Member", nil, nil)
		if err != nil {
			return err
		}
	} else if !constants.IsRosterFacility(controller.Facility) {
		err := core.AddHold(controller, constants.RCEExam, "Automatic", nil, nil)
		if err != nil {
			return err
		}
	}

	// Promotion
	var ratingChange *database.RatingChange
	for _, rc := range controller.RatingChanges {
		if rc.FromRating > constants.C1 {
			continue // Skip post-C1 promotions
		}
		if ratingChange == nil || rc.CreatedAt.After(ratingChange.CreatedAt) {
			ratingChange = &rc
		}
	}
	if ratingChange != nil && ratingChange.CreatedAt.AddDate(0, 0, 90).After(time.Now()) {
		expiresAt := ratingChange.CreatedAt.AddDate(0, 0, 90)
		err := core.AddHold(controller, constants.RecentPromotion, "Recent Promotion", &expiresAt, nil)
		if err != nil {
			return err
		}
	}

	// Transfer TODO

	// PendingTransfer TODO
	return nil
}
