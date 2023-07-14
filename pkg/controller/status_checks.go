package controller

import (
	db "github.com/VATUSA/api-v3/pkg/database"
	"github.com/VATUSA/api-v3/pkg/hold"
	"time"
)

func IsTransferEligible(c *db.Controller) bool {
	if !c.IsInDivision {
		return false
	}
	if !c.IsActive {
		return false
	}
	for _, h := range c.Holds {
		if h.ExpiresAt.Before(time.Now()) {
			db.DB.Delete(&h)
		} else {
			holdMeta := hold.Get(h.Hold)
			if holdMeta.PreventTransfer {
				return false
			}
		}
	}
	return true
}

func IsVisitEligible(c *db.Controller) bool {
	if !c.IsInDivision && !c.IsApprovedExternalVisitor {
		return false
	}
	if !c.IsActive {
		return false
	}
	for _, h := range c.Holds {
		if h.ExpiresAt.Before(time.Now()) {
			db.DB.Delete(&h)
		} else {
			holdMeta := hold.Get(h.Hold)
			if holdMeta.PreventVisit {
				return false
			}
		}
	}
	return true
}

func IsPromotionEligible(c *db.Controller) bool {
	if !c.IsInDivision {
		return false
	}
	if !c.IsActive {
		return false
	}
	for _, h := range c.Holds {
		if h.ExpiresAt.Before(time.Now()) {
			db.DB.Delete(&h)
		} else {
			holdMeta := hold.Get(h.Hold)
			if holdMeta.PreventPromotion {
				return false
			}
		}
	}
	// TODO
	return true
}
