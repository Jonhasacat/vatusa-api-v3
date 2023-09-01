package translator

import (
	"github.com/VATUSA/api-v3/internal/database"
	"github.com/VATUSA/api-v3/pkg/self_api/model"
)

func TranslateRatingChange(change database.RatingChange) model.RatingChange {
	return model.RatingChange{
		FromRating:  TranslateRating(change.FromRating),
		ToRating:    TranslateRating(change.ToRating),
		CreatedAt:   change.CreatedAt,
		RequesterID: change.AdminID,
	}
}

func TranslateRatingChanges(changes []database.RatingChange) []model.RatingChange {
	out := make([]model.RatingChange, 0)
	for _, change := range changes {
		out = append(out, TranslateRatingChange(change))
	}
	return out
}
