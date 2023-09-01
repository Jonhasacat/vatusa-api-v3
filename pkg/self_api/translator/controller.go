package translator

import (
	"github.com/VATUSA/api-v3/internal/core"
	db "github.com/VATUSA/api-v3/internal/database"
	"github.com/VATUSA/api-v3/pkg/self_api/model"
)

func TranslateController(controller db.Controller) model.Controller {
	return model.Controller{
		ID:                  controller.Id,
		DisplayName:         core.DisplayName(&controller),
		VATSIMRating:        TranslateRating(controller.Certificate.Rating),
		ATCRating:           TranslateRating(controller.ATCRating),
		Facility:            controller.Facility,
		FacilityJoin:        controller.FacilityJoin,
		LastPromotion:       controller.LastPromotion,
		InDivision:          controller.IsInDivision,
		IsTransferEligible:  core.IsTransferEligible(&controller),
		IsVisitEligible:     core.IsVisitEligible(&controller),
		IsPromotionEligible: core.IsPromotionEligible(&controller),
		DiscordId:           controller.DiscordId,
		RatingChanges:       TranslateRatingChanges(controller.RatingChanges),
		Transfers:           TranslateTransfers(controller.Transfers),
		Visits:              TranslateVisits(controller.Visits),
		Roles:               TranslateRoles(controller.Roles),
	}
}

func TranslateControllers(controllers []db.Controller) []model.Controller {
	out := make([]model.Controller, 0)
	for _, controller := range controllers {
		out = append(out, TranslateController(controller))

	}
	return out
}
