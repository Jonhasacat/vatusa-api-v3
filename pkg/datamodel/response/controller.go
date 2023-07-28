package response

import (
	"github.com/VATUSA/api-v3/internal/core"
	"github.com/VATUSA/api-v3/pkg/constants"
	"github.com/VATUSA/api-v3/pkg/database"
	"time"
)

type Controller struct {
	CertificateID       uint64
	CertificateName     *string
	DisplayName         string
	Email               *string
	VATSIMRating        ControllerRating
	ATCRating           ControllerRating
	Facility            string
	FacilityJoin        *time.Time
	LastPromotion       *time.Time
	InDivision          bool
	IsTransferEligible  bool
	IsVisitEligible     bool
	IsPromotionEligible bool
	DiscordId           *string
	RatingChanges       []ControllerRatingChange
	Transfers           []ControllerTransfer
	Visits              []string
	Roles               []ControllerRole
}

func MakeController(c *database.Controller) *Controller {
	certificateName := core.CertificateName(c)
	model := &Controller{
		CertificateID:   c.CertificateId,
		CertificateName: &certificateName,
		DisplayName:     core.DisplayName(c),
		Email:           &c.Certificate.Email,
		VATSIMRating: ControllerRating{
			Value: c.Certificate.Rating,
			Short: constants.ShortMap[constants.Rating(c.Certificate.Rating)],
			Long:  constants.LongMap[constants.Rating(c.Certificate.Rating)],
		},
		ATCRating: ControllerRating{
			Value: c.ATCRating,
			Short: constants.ShortMap[constants.Rating(c.ATCRating)],
			Long:  constants.LongMap[constants.Rating(c.ATCRating)],
		},
		Facility:            c.Facility,
		FacilityJoin:        c.FacilityJoin,
		LastPromotion:       c.LastPromotion,
		InDivision:          c.IsInDivision,
		IsTransferEligible:  core.IsTransferEligible(c),
		IsVisitEligible:     core.IsVisitEligible(c),
		IsPromotionEligible: core.IsPromotionEligible(c),
		DiscordId:           c.DiscordId,
		RatingChanges:       []ControllerRatingChange{},
		Transfers:           []ControllerTransfer{},
		Visits:              []string{},
		Roles:               []ControllerRole{},
	}
	for _, r := range c.Roles {
		model.Roles = append(model.Roles, *MakeControllerRoleResponse(&r))
	}
	for _, v := range c.Visits {
		model.Visits = append(model.Visits, v.Facility)
	}
	for _, rc := range c.RatingChanges {
		model.RatingChanges = append(model.RatingChanges, *MakeControllerRatingChange(&rc))
	}
	for _, t := range c.Transfers {
		model.Transfers = append(model.Transfers, *MakeControllerTransfer(&t))
	}
	return model
}

func MakeControllerRedacted(c *database.Controller) *Controller {
	model := MakeController(c)
	model.CertificateName = nil
	model.Email = nil
	return model
}
