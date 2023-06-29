package model

import "vatusa-api-v3/database"

type ControllerTransfer struct {
	FromFacility string
	ToFacility   string
	Reason       string
}

func MakeControllerTransfer(t *database.Transfer) *ControllerTransfer {
	transfer := &ControllerTransfer{
		FromFacility: string(t.FromFacility),
		ToFacility:   string(t.ToFacility),
		Reason:       t.Reason,
	}
	return transfer
}
