package core

import (
	"errors"
	"fmt"
	database2 "github.com/VATUSA/api-v3/internal/database"
	"github.com/VATUSA/api-v3/pkg/constants"
	"time"
)

func AcceptRosterRequest(request *database2.ControllerRosterRequest, reason *string, requester *database2.Controller) error {
	if request.Status != constants.StatusPending {
		return errors.New(fmt.Sprintf("ControllerRosterRequest %d is not pending", request.ID))
	}
	now := time.Now()
	request.Status = constants.StatusAccepted
	request.StatusReason = reason
	request.AdminID = requester.Id

	transfer := database2.Transfer{
		ControllerID: request.ControllerID,
		Controller:   request.Controller,
		FromFacility: request.Controller.Facility,
		ToFacility:   request.Facility,
		Reason:       request.Reason,
	}

	request.Controller.Facility = request.Facility
	request.Controller.FacilityJoin = &now
	database2.DB.Save(request)
	database2.DB.Save(transfer)
	request.Controller.Save()
	// TODO
	return nil
}

func RejectRosterRequest(request *database2.ControllerRosterRequest, reason *string, requester *database2.Controller) error {
	if request.Status != constants.StatusPending {
		return errors.New(fmt.Sprintf("ControllerRosterRequest %d is not pending", request.ID))
	}
	if reason == nil {
		return errors.New("reason is required to reject a request")
	}
	request.Status = constants.StatusAccepted
	request.StatusReason = reason
	request.AdminID = requester.Id
	database2.DB.Save(request)
	// TODO
	return nil
}
