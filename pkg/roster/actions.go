package roster

import (
	"errors"
	"fmt"
	"github.com/VATUSA/api-v3/pkg/action_log"
	"github.com/VATUSA/api-v3/pkg/controller"
	"github.com/VATUSA/api-v3/pkg/database"
	"github.com/VATUSA/api-v3/pkg/facility"
	"github.com/VATUSA/api-v3/pkg/role"
)

func RequestTransfer(c *database.Controller, fac facility.Facility, reason string) error {
	if !controller.IsTransferEligible(c) {
		return errors.New(fmt.Sprintf("Controller %d is not eligible for transfer", c.CertificateId))
	}
	if !facility.IsRosterFacility(fac) {
		return errors.New(fmt.Sprintf("Facility %s is not eligible for transfer requests", fac))
	}
	// TODO: Create transfer request
	// TODO: Send notification
	return nil
}

func ForceTransfer(controller *database.Controller, fac facility.Facility, reason string) error {
	if controller.Facility == fac {
		return errors.New(fmt.Sprintf("controller %d is already in facility %s", controller.Id, fac))
	}
	transfer := &database.Transfer{
		ControllerID: controller.Id,
		Controller:   controller,
		FromFacility: controller.Facility,
		ToFacility:   fac,
		Reason:       reason,
	}
	transfer.Save()
	return nil
}

func RemoveFromFacility(controller *database.Controller, requester *database.Controller, reason string) error {
	if controller.Facility == facility.Academy ||
		controller.Facility == facility.NonMember ||
		controller.Facility == facility.Inactive {
		return errors.New(fmt.Sprintf("Cannot remove a controller from this facility"))
	}
	message := fmt.Sprintf("Removed from %s: %s", controller.Facility, reason)
	oldFacility := controller.Facility
	err := ForceTransfer(controller, "ZAE", message)
	if err != nil {
		return err
	}
	err = action_log.LogAction(controller, message, requester)
	if err != nil {
		return err
	}
	err = role.RemoveFacilityRoles(controller, oldFacility)
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

func AddVisitor(c *database.Controller, fac facility.Facility, requester *database.Controller, reason string) error {
	if c.Facility == facility.Academy || c.Facility == facility.Inactive {
		return errors.New(fmt.Sprintf(
			"controller %d is in %s and can not visit", c.Id, c.Facility))
	}
	if !controller.IsVisitEligible(c) {
		return errors.New(fmt.Sprintf("controller %d is not eligible to visit", c.Id))
	}
	if !facility.IsRosterFacility(fac) {
		return errors.New(fmt.Sprintf("facility %s is not eligible for visitors", fac))
	}
	if IsVisiting(c, fac) {
		return errors.New(fmt.Sprintf("controller %d is already visiting %s", c.Id, fac))
	}
	visit := database.ControllerVisit{
		ControllerID: c.Id,
		Controller:   c,
		Facility:     fac,
	}
	visit.Save()
	err := action_log.LogAction(c, fmt.Sprintf("Added to %s visitor roster: %s", fac, reason), requester)
	if err != nil {
		return err
	}
	// TODO: Send notification
	return nil
}

func RemoveVisitor(controller *database.Controller, fac facility.Facility, requester *database.Controller, reason string) error {
	for _, v := range controller.Visits {
		if v.Facility == fac {
			result := database.DB.Delete(v)
			if result.Error != nil {
				return result.Error
			}
			err := role.RemoveFacilityRoles(controller, fac)
			if err != nil {
				return err
			}
			err = action_log.LogAction(controller, fmt.Sprintf("Removed from %s visitor roster: %s", fac, reason), requester)
			if err != nil {
				return err
			}
			err = role.RemoveFacilityRoles(controller, fac)
			if err != nil {
				return err
			}
			// TODO: Send notification
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Controller %d is not visiting %s", controller.Id, fac))
}

func RemoveAllVisits(controller *database.Controller, requester *database.Controller, reason string) error {
	for _, v := range controller.Visits {
		result := database.DB.Delete(v)
		if result.Error != nil {
			return result.Error
		}
		err := role.RemoveFacilityRoles(controller, v.Facility)
		if err != nil {
			return err
		}
		err = action_log.LogAction(controller, fmt.Sprintf("Removed from %s visitor roster: %s", v.Facility, reason), requester)
		if err != nil {
			return err
		}
		err = role.RemoveFacilityRoles(controller, v.Facility)
		if err != nil {
			return err
		}
		// TODO: Send notification
	}
	return nil
}
