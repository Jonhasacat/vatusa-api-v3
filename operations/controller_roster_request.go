package operations

import (
	"errors"
	"fmt"
	"time"
	"vatusa-api-v3/constants"
	db "vatusa-api-v3/database"
)

func AcceptRosterRequest(request *db.ControllerRosterRequest, reason *string, requester *db.Controller) error {
	if request.Status != constants.RequestPending {
		return errors.New(fmt.Sprintf("ControllerRosterRequest %d is not pending", request.ID))
	}
	now := time.Now()
	request.Status = constants.RequestAccepted
	request.StatusReason = reason
	request.AdminID = requester.Id

	transfer := db.Transfer{
		ControllerID: request.ControllerID,
		Controller:   request.Controller,
		FromFacility: request.Controller.Facility,
		ToFacility:   request.Facility,
		Reason:       request.Reason,
	}

	request.Controller.Facility = request.Facility
	request.Controller.FacilityJoin = &now
	db.DB.Save(request)
	db.DB.Save(transfer)
	request.Controller.Save()
	// TODO
	return nil
}

func RejectRosterRequest(request *db.ControllerRosterRequest, reason *string, requester *db.Controller) error {
	if request.Status != constants.RequestPending {
		return errors.New(fmt.Sprintf("ControllerRosterRequest %d is not pending", request.ID))
	}
	if reason == nil {
		return errors.New("reason is required to reject a request")
	}
	request.Status = constants.RequestAccepted
	request.StatusReason = reason
	request.AdminID = requester.Id
	db.DB.Save(request)
	// TODO
	return nil
}
