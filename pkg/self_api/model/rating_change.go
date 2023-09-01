package model

import "time"

type RatingChange struct {
	FromRating  Rating
	ToRating    Rating
	CreatedAt   time.Time
	RequesterID uint64
}
