package model

import (
	"time"
	"vatusa-api-v3/constants"
	"vatusa-api-v3/database"
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
			Value: int(rc.FromRating),
			Short: constants.RatingShortMap[rc.FromRating],
			Long:  constants.RatingLongMap[rc.FromRating],
		},
		ToRating: ControllerRating{
			Value: int(rc.ToRating),
			Short: constants.RatingShortMap[rc.ToRating],
			Long:  constants.RatingLongMap[rc.ToRating],
		},
		CreatedAt:    time.Time{},
		RequesterCID: 0,
	}
	return ratingChange
}
