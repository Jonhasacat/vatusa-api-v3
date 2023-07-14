package response

import (
	"github.com/VATUSA/api-v3/pkg/controller"
	"github.com/VATUSA/api-v3/pkg/database"
	"github.com/VATUSA/api-v3/pkg/rating"
	"time"
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
	model := &ControllerLimited{
		CertificateID: c.CertificateId,
		DisplayName:   controller.DisplayName(c),
		VATSIMRating: ControllerRating{
			Value: c.Certificate.Rating,
			Short: rating.ShortMap[rating.Rating(c.Certificate.Rating)],
			Long:  rating.LongMap[rating.Rating(c.Certificate.Rating)],
		},
		ATCRating: ControllerRating{
			Value: c.ATCRating,
			Short: rating.ShortMap[rating.Rating(c.ATCRating)],
			Long:  rating.LongMap[rating.Rating(c.ATCRating)],
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
		model.Roles = append(model.Roles, *MakeControllerRoleResponse(&v))
	}
	for _, v := range c.Visits {
		model.Visits = append(model.Visits, v.Facility)
	}
	return model
}
