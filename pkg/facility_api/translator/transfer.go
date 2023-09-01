package translator

import (
	"github.com/VATUSA/api-v3/internal/database"
	"github.com/VATUSA/api-v3/pkg/facility_api/model"
)

func TranslateTransfer(transfer database.Transfer) model.Transfer {
	return model.Transfer{
		FromFacility: transfer.FromFacility,
		ToFacility:   transfer.ToFacility,
		Reason:       transfer.Reason,
		DateTime:     transfer.CreatedAt,
	}
}

func TranslateTransfers(transfers []database.Transfer) []model.Transfer {
	out := make([]model.Transfer, 0)
	for _, transfer := range transfers {
		out = append(out, TranslateTransfer(transfer))
	}
	return out
}
