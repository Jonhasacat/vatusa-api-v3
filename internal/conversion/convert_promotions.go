package conversion

import (
	"fmt"
	"github.com/VATUSA/api-v3/internal/conversion/legacydb"
	db "github.com/VATUSA/api-v3/pkg/database"
)

func LoadLegacyPromotions() ([]legacydb.Promotion, error) {
	var promotions []legacydb.Promotion
	result := legacydb.DB.Model(legacydb.Promotion{}).Find(&promotions)
	if result.Error != nil {
		return nil, result.Error
	}
	return promotions, nil
}

func LoadLegacyPromotionsByCID(cid uint64) ([]legacydb.Promotion, error) {
	var promotions []legacydb.Promotion
	result := legacydb.DB.Model(legacydb.Promotion{}).Where("cid = ?", cid).Find(&promotions)
	if result.Error != nil {
		return nil, result.Error
	}
	return promotions, nil
}

func ProcessLegacyPromotion(promotion legacydb.Promotion) error {
	c, err := db.FetchControllerByCID(promotion.CID)
	if err != nil {
		return err
	}
	record := &db.RatingChange{
		ControllerID: c.Id,
		Controller:   c,
		FromRating:   promotion.FromRating,
		ToRating:     promotion.ToRating,
		AdminID:      promotion.Grantor,
	}
	result := db.DB.Save(record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func ConvertPromotions() error {
	promotions, err := LoadLegacyPromotions()
	if err != nil {
		return err
	}
	for _, p := range promotions {
		err = ProcessLegacyPromotion(p)
		if err != nil {
			println(fmt.Sprintf("Error Converting Promotion %d: %s", p.ID, err.Error()))
		}
	}
	return nil
}
