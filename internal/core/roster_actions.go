package core

import (
	"errors"
	"fmt"
	database2 "github.com/VATUSA/api-v3/internal/database"
	"github.com/VATUSA/api-v3/pkg/constants"
)

func RequestTransfer(c *database2.Controller, fac constants.Facility, reason string) error {
	if !IsTransferEligible(c) {
		return errors.New(fmt.Sprintf("Controller %d is not eligible for transfer", c.CertificateId))
	}
	if !constants.IsRosterFacility(fac) {
		return errors.New(fmt.Sprintf("Facility %s is not eligible for transfer requests", fac))
	}
	// TODO: Create transfer request
	// TODO: Send notification
	return nil
}

func ForceTransfer(controller *database2.Controller, fac constants.Facility, reason string) error {
	if controller.Facility == fac {
		return errors.New(fmt.Sprintf("controller %d is already in facility %s", controller.Id, fac))
	}
	transfer := &database2.Transfer{
		ControllerID: controller.Id,
		Controller:   controller,
		FromFacility: controller.Facility,
		ToFacility:   fac,
		Reason:       reason,
	}
	transfer.Save()
	return nil
}

func RemoveFromFacility(controller *database2.Controller, requester *database2.Controller, reason string) error {
	if controller.Facility == constants.Academy ||
		controller.Facility == constants.NonMember ||
		controller.Facility == constants.InactiveFacility {
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

func AddVisitor(c *database2.Controller, fac constants.Facility, requester *database2.Controller, reason string) error {
	if c.Facility == constants.Academy || c.Facility == constants.InactiveFacility {
		return errors.New(fmt.Sprintf(
			"controller %d is in %s and can not visit", c.Id, c.Facility))
	}
	if !IsVisitEligible(c) {
		return errors.New(fmt.Sprintf("controller %d is not eligible to visit", c.Id))
	}
	if !constants.IsRosterFacility(fac) {
		return errors.New(fmt.Sprintf("facility %s is not eligible for visitors", fac))
	}
	if IsVisiting(c, fac) {
		return errors.New(fmt.Sprintf("controller %d is already visiting %s", c.Id, fac))
	}
	visit := database2.ControllerVisit{
		ControllerID: c.Id,
		Controller:   c,
		Facility:     fac,
	}
	visit.Save()
	err := LogAction(c, fmt.Sprintf("Added to %s visitor roster: %s", fac, reason), requester)
	if err != nil {
		return err
	}
	// TODO: Send notification
	return nil
}

func RemoveVisitor(controller *database2.Controller, fac constants.Facility, requester *database2.Controller, reason string) error {
	for _, v := range controller.Visits {
		if v.Facility == fac {
			result := database2.DB.Delete(v)
			if result.Error != nil {
				return result.Error
			}
			err := RemoveFacilityRoles(controller, fac)
			if err != nil {
				return err
			}
			err = LogAction(controller, fmt.Sprintf("Removed from %s visitor roster: %s", fac, reason), requester)
			if err != nil {
				return err
			}
			err = RemoveFacilityRoles(controller, fac)
			if err != nil {
				return err
			}
			// TODO: Send notification
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Controller %d is not visiting %s", controller.Id, fac))
}

func RemoveAllVisits(controller *database2.Controller, requester *database2.Controller, reason string) error {
	for _, v := range controller.Visits {
		result := database2.DB.Delete(v)
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
