package core

import (
	"github.com/VATUSA/api-v3/internal/database"
	"github.com/VATUSA/api-v3/pkg/constants"
	"time"
)

func IsTransferEligible(c *database.Controller) bool {
	if !c.IsInDivision {
		return false
	}
	if !c.IsActive {
		return false
	}
	for _, h := range c.Holds {
		if h.ExpiresAt.Before(time.Now()) {
			database.DB.Delete(&h)
		} else {
			holdMeta := constants.Get(h.Hold)
			if holdMeta.PreventTransfer {
				return false
			}
		}
	}
	return true
}

func IsVisitEligible(c *database.Controller) bool {
	if !c.IsInDivision && !c.IsApprovedExternalVisitor {
		return false
	}
	if !c.IsActive {
		return false
	}
	for _, h := range c.Holds {
		if h.ExpiresAt.Before(time.Now()) {
			database.DB.Delete(&h)
		} else {
			holdMeta := constants.Get(h.Hold)
			if holdMeta.PreventVisit {
				return false
			}
		}
	}
	return true
}

func IsPromotionEligible(c *database.Controller) bool {
	if !c.IsInDivision {
		return false
	}
	if !c.IsActive {
		return false
	}
	if c.ATCRating >= constants.C1 {
		return false
	}
	for _, h := range c.Holds {
		if h.ExpiresAt.Before(time.Now()) {
			database.DB.Delete(&h)
		} else {
			holdMeta := constants.Get(h.Hold)
			if holdMeta.PreventPromotion {
				return false
			}
		}
	}
	// TODO
	return true
}
