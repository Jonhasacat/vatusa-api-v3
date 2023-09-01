package model

import "time"

type Controller struct {
	ID                  uint64
	DisplayName         string
	Email               string
	VATSIMRating        Rating
	ATCRating           Rating
	Facility            string
	FacilityJoin        *time.Time
	LastPromotion       *time.Time
	InDivision          bool
	IsTransferEligible  bool
	IsVisitEligible     bool
	IsPromotionEligible bool
	DiscordId           *string
	RatingChanges       []RatingChange
	Transfers           []Transfer
	Visits              []string
	Roles               []Role
}
