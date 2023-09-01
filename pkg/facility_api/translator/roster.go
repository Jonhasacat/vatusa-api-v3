package translator

import (
	db "github.com/VATUSA/api-v3/internal/database"
	"github.com/VATUSA/api-v3/pkg/facility_api/model"
)

func TranslateRosterRequest(request db.ControllerRosterRequest) model.RosterRequest {
	return model.RosterRequest{
		ID:         request.ID,
		Controller: TranslateController(*request.Controller),
		Type:       model.RosterRequestType(request.RequestType),
		Reason:     request.Reason,
	}
}

func TranslateRosterRequests(requests []db.ControllerRosterRequest) []model.RosterRequest {
	out := make([]model.RosterRequest, 0)
	for _, request := range requests {
		out = append(out, TranslateRosterRequest(request))
	}
	return out
}
