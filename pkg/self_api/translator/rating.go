package translator

import (
	"github.com/VATUSA/api-v3/pkg/constants"
	"github.com/VATUSA/api-v3/pkg/self_api/model"
)

func TranslateRating(value int) model.Rating {
	return model.Rating{
		Value: value,
		Short: constants.RatingShort(value),
		Long:  constants.RatingLong(value),
	}
}
