package model

import (
	"time"
	"vatusa-api-v3/constants"
	"vatusa-api-v3/database"
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
	certificateName := c.CertificateName()
	controller := &Controller{
		CertificateID:   c.CertificateId,
		CertificateName: &certificateName,
		DisplayName:     c.DisplayName(),
		Email:           &c.Certificate.Email,
		VATSIMRating: ControllerRating{
			Value: int(c.Certificate.Rating),
			Short: constants.RatingShortMap[c.Certificate.Rating],
			Long:  constants.RatingLongMap[c.Certificate.Rating],
		},
		ATCRating: ControllerRating{
			Value: int(c.ATCRating),
			Short: constants.RatingShortMap[c.ATCRating],
			Long:  constants.RatingLongMap[c.ATCRating],
		},
		Facility:            string(c.Facility),
		FacilityJoin:        c.FacilityJoin,
		LastPromotion:       c.LastPromotion,
		InDivision:          c.IsInDivision,
		IsTransferEligible:  c.IsTransferEligible(),
		IsVisitEligible:     c.IsVisitEligible(),
		IsPromotionEligible: c.IsPromotionEligible(),
		DiscordId:           c.DiscordId,
		RatingChanges:       []ControllerRatingChange{},
		Transfers:           []ControllerTransfer{},
		Visits:              []string{},
		Roles:               []ControllerRole{},
	}
	for _, r := range c.Roles {
		controller.Roles = append(controller.Roles, *MakeControllerRoleResponse(&r))
	}
	for _, v := range c.Visits {
		controller.Visits = append(controller.Visits, string(v.Facility))
	}
	for _, rc := range c.RatingChanges {
		controller.RatingChanges = append(controller.RatingChanges, *MakeControllerRatingChange(&rc))
	}
	for _, t := range c.Transfers {
		controller.Transfers = append(controller.Transfers, *MakeControllerTransfer(&t))
	}
	return controller
}

func MakeControllerRedacted(c *database.Controller) *Controller {
	controller := MakeController(c)
	controller.CertificateName = nil
	controller.Email = nil
	return controller
}
