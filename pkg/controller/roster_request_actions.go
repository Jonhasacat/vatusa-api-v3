package controller

import (
	"errors"
	"fmt"
	"github.com/VATUSA/api-v3/pkg/database"
	"github.com/VATUSA/api-v3/pkg/roster"
	"time"
)

func AcceptRosterRequest(request *database.ControllerRosterRequest, reason *string, requester *database.Controller) error {
	if request.Status != roster.StatusPending {
		return errors.New(fmt.Sprintf("ControllerRosterRequest %d is not pending", request.ID))
	}
	now := time.Now()
	request.Status = roster.StatusAccepted
	request.StatusReason = reason
	request.AdminID = requester.Id

	transfer := database.Transfer{
		ControllerID: request.ControllerID,
		Controller:   request.Controller,
		FromFacility: request.Controller.Facility,
		ToFacility:   request.Facility,
		Reason:       request.Reason,
	}

	request.Controller.Facility = request.Facility
	request.Controller.FacilityJoin = &now
	database.DB.Save(request)
	database.DB.Save(transfer)
	request.Controller.Save()
	// TODO
	return nil
}

func RejectRosterRequest(request *database.ControllerRosterRequest, reason *string, requester *database.Controller) error {
	if request.Status != roster.StatusPending {
		return errors.New(fmt.Sprintf("ControllerRosterRequest %d is not pending", request.ID))
	}
	if reason == nil {
		return errors.New("reason is required to reject a request")
	}
	request.Status = roster.StatusAccepted
	request.StatusReason = reason
	request.AdminID = requester.Id
	database.DB.Save(request)
	// TODO
	return nil
}
