package response

import (
	"github.com/VATUSA/api-v3/internal/database"
)

type ControllerTransfer struct {
	FromFacility string
	ToFacility   string
	Reason       string
}

func MakeControllerTransfer(t *database.Transfer) *ControllerTransfer {
	transfer := &ControllerTransfer{
		FromFacility: t.FromFacility,
		ToFacility:   t.ToFacility,
		Reason:       t.Reason,
	}
	return transfer
}
