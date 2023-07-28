package response

import (
	"github.com/VATUSA/api-v3/internal/core"
	"github.com/VATUSA/api-v3/pkg/constants"
	"github.com/VATUSA/api-v3/pkg/database"
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
		DisplayName:   core.DisplayName(c),
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
