package translator

import (
	db "github.com/VATUSA/api-v3/internal/database"
)

func TranslateVisit(visit db.ControllerVisit) string {
	return visit.Facility
}

func TranslateVisits(visits []db.ControllerVisit) []string {
	out := make([]string, 0)
	for _, visit := range visits {
		out = append(out, TranslateVisit(visit))
	}
	return out
}
