package operations

import (
	"errors"
	"fmt"
	"vatusa-api-v3/constants"
	db "vatusa-api-v3/database"
)

func RequestTransfer(controller *db.Controller, facility constants.Facility, reason string) error {
	if !controller.IsTransferEligible() {
		return errors.New(fmt.Sprintf("Controller %d is not eligible for transfer", controller.CertificateId))
	}
	if !constants.IsRosterFacility(facility) {
		return errors.New(fmt.Sprintf("Facility %s is not eligible for transfer requests", facility))
	}
	// TODO: Create transfer request
	// TODO: Send notification
	return nil
}

func ForceTransfer(controller *db.Controller, facility constants.Facility, reason string) error {
	if controller.Facility == facility {
		return errors.New(fmt.Sprintf("controller %d is already in facility %s", controller.Id, facility))
	}
	transfer := &db.Transfer{
		ControllerID: controller.Id,
		Controller:   controller,
		FromFacility: controller.Facility,
		ToFacility:   facility,
		Reason:       reason,
	}
	transfer.Save()
	return nil
}

func RemoveFromFacility(controller *db.Controller, requester *db.Controller, reason string) error {
	if controller.Facility == constants.Academy ||
		controller.Facility == constants.NonMember ||
		controller.Facility == constants.Inactive {
		return errors.New(fmt.Sprintf("Cannot remove a controller from this facility"))
	}
	message := fmt.Sprintf("Removed from %s: %s", controller.Facility, reason)
	oldFacility := controller.Facility
	err := ForceTransfer(controller, "ZAE", message)
	if err != nil {
		return err
	}
	err = LogAction(controller, message, requester)
	if err != nil {
		return err
	}
	err = RemoveFacilityRoles(controller, oldFacility)
	if err != nil {
		return err
	}
	err = RemoveAllVisits(controller, nil, "Transfer to ZAE")
	if err != nil {
		return err
	}
	// TODO: Send notification
	return nil
}

func AddVisitor(controller *db.Controller, facility constants.Facility, requester *db.Controller, reason string) error {
	if controller.Facility == constants.Academy || controller.Facility == constants.Inactive {
		return errors.New(fmt.Sprintf(
			"controller %d is in %s and can not visit", controller.Id, controller.Facility))
	}
	if !controller.IsVisitEligible() {
		return errors.New(fmt.Sprintf("controller %d is not eligible to visit", controller.Id))
	}
	if !constants.IsRosterFacility(facility) {
		return errors.New(fmt.Sprintf("facility %s is not eligible for visitors", facility))
	}
	if controller.IsVisiting(facility) {
		return errors.New(fmt.Sprintf("controller %d is already visiting %s", controller.Id, facility))
	}
	visit := db.ControllerVisit{
		ControllerID: controller.Id,
		Controller:   controller,
		Facility:     facility,
	}
	visit.Save()
	err := LogAction(controller, fmt.Sprintf("Added to %s visitor roster: %s", facility, reason), requester)
	if err != nil {
		return err
	}
	// TODO: Send notification
	return nil
}

func RemoveVisitor(controller *db.Controller, facility constants.Facility, requester *db.Controller, reason string) error {
	for _, v := range controller.Visits {
		if v.Facility == facility {
			result := db.DB.Delete(v)
			if result.Error != nil {
				return result.Error
			}
			err := RemoveFacilityRoles(controller, facility)
			if err != nil {
				return err
			}
			err = LogAction(controller, fmt.Sprintf("Removed from %s visitor roster: %s", facility, reason), requester)
			if err != nil {
				return err
			}
			err = RemoveFacilityRoles(controller, facility)
			if err != nil {
				return err
			}
			// TODO: Send notification
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Controller %d is not visiting %s", controller.Id, facility))
}

func RemoveAllVisits(controller *db.Controller, requester *db.Controller, reason string) error {
	for _, v := range controller.Visits {
		result := db.DB.Delete(v)
		if result.Error != nil {
			return result.Error
		}
		err := RemoveFacilityRoles(controller, v.Facility)
		if err != nil {
			return err
		}
		err = LogAction(controller, fmt.Sprintf("Removed from %s visitor roster: %s", v.Facility, reason), requester)
		if err != nil {
			return err
		}
		err = RemoveFacilityRoles(controller, v.Facility)
		if err != nil {
			return err
		}
		// TODO: Send notification
	}
	return nil
}
