package response

import (
	"github.com/VATUSA/api-v3/pkg/database"
)

type ControllerRosterRequest struct {
	ID           uint
	Controller   *Controller
	Facility     string
	RequestType  string
	Reason       string
	RequesterID  uint64
	Requester    *Controller
	Status       string
	StatusReason *string
}

func MakeControllerRosterRequestResponse(r *database.ControllerRosterRequest) *ControllerRosterRequest {
	return &ControllerRosterRequest{
		ID:           r.ID,
		Controller:   MakeController(r.Controller),
		Facility:     string(r.Facility),
		RequestType:  r.RequestType,
		Reason:       r.Reason,
		RequesterID:  r.RequesterID,
		Requester:    MakeController(r.Controller),
		Status:       r.Status,
		StatusReason: r.StatusReason,
	}
}
