package response

import (
	"github.com/VATUSA/api-v3/internal/database"
	"github.com/VATUSA/api-v3/pkg/constants"
	"time"
)

type ControllerRatingChange struct {
	FromRating   ControllerRating
	ToRating     ControllerRating
	CreatedAt    time.Time
	RequesterCID uint64
}

func MakeControllerRatingChange(rc *database.RatingChange) *ControllerRatingChange {
	ratingChange := &ControllerRatingChange{
		FromRating: ControllerRating{
			Value: rc.FromRating,
			Short: constants.RatingShort(rc.FromRating),
			Long:  constants.RatingLong(rc.FromRating),
		},
		ToRating: ControllerRating{
			Value: rc.ToRating,
			Short: constants.RatingShort(rc.ToRating),
			Long:  constants.RatingLong(rc.ToRating),
		},
		CreatedAt:    time.Time{},
		RequesterCID: 0,
	}
	return ratingChange
}
