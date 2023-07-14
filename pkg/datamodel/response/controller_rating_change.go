package response

import (
	"github.com/VATUSA/api-v3/pkg/database"
	"github.com/VATUSA/api-v3/pkg/rating"
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
			Short: rating.ShortFromInt(rc.FromRating),
			Long:  rating.LongFromInt(rc.FromRating),
		},
		ToRating: ControllerRating{
			Value: rc.ToRating,
			Short: rating.ShortFromInt(rc.ToRating),
			Long:  rating.LongFromInt(rc.ToRating),
		},
		CreatedAt:    time.Time{},
		RequesterCID: 0,
	}
	return ratingChange
}
