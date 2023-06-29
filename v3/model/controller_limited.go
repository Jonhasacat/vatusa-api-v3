package model

import (
	"time"
	"vatusa-api-v3/constants"
	"vatusa-api-v3/database"
)

type ControllerLimited struct {
	CertificateID uint64
	DisplayName   string
	VATSIMRating  ControllerRating
	ATCRating     ControllerRating
	Facility      string
	FacilityJoin  *time.Time
	LastPromotion *time.Time
	InDivision    bool
	DiscordId     *string
	Visits        []string
	Roles         []ControllerRole
}

func MakeControllerLimited(c *database.Controller) *ControllerLimited {
	controller := &ControllerLimited{
		CertificateID: c.CertificateId,
		DisplayName:   c.DisplayName(),
		VATSIMRating: ControllerRating{
			Value: c.Certificate.Rating,
			Short: constants.RatingShortMap[c.Certificate.Rating],
			Long:  constants.RatingLongMap[c.Certificate.Rating],
		},
		ATCRating: ControllerRating{
			Value: c.ATCRating,
			Short: constants.RatingShortMap[c.ATCRating],
			Long:  constants.RatingLongMap[c.ATCRating],
		},
		Facility:      c.Facility,
		FacilityJoin:  c.FacilityJoin,
		LastPromotion: c.LastPromotion,
		InDivision:    c.IsInDivision,
		DiscordId:     c.DiscordId,
		Visits:        []string{},
		Roles:         []ControllerRole{},
	}
	for _, v := range c.Roles {
		controller.Roles = append(controller.Roles, MakeControllerRoleResponse(&v))
	}
	for _, v := range c.Visits {
		controller.Visits = append(controller.Visits, v.Facility)
	}
	return controller
}
