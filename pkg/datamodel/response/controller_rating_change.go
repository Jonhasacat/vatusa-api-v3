package response

import (
	"github.com/VATUSA/api-v3/pkg/constants"
	"github.com/VATUSA/api-v3/pkg/database"
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
			Short: constants.ShortFromInt(rc.FromRating),
			Long:  constants.LongFromInt(rc.FromRating),
		},
		ToRating: ControllerRating{
			Value: rc.ToRating,
			Short: constants.ShortFromInt(rc.ToRating),
			Long:  constants.LongFromInt(rc.ToRating),
		},
		CreatedAt:    time.Time{},
		RequesterCID: 0,
	}
	return ratingChange
}
