package core

import (
	"errors"
	"fmt"
	"github.com/VATUSA/api-v3/internal/database"
	"github.com/VATUSA/api-v3/pkg/constants"
	"time"
)

func AddHold(controller *database.Controller, hold constants.Hold, reason string, expiresAt *time.Time, requester *database.Controller) error {
	if HasHold(controller, hold) {
		return nil
	}
	record := database.ControllerHold{
		Controller: controller,
		Hold:       hold,
		ExpiresAt:  expiresAt,
	}
	err := record.Save()
	if err != nil {
		return err
	}
	err = LogAction(controller, fmt.Sprintf("Hold %s applied: %s", hold, reason), requester)
	return nil
}

func HasHold(controller *database.Controller, hold constants.Hold) bool {
	for _, h := range controller.Holds {
		if h.Hold == hold {
			return true
		}
	}
	return false
}

func RemoveHold(controller *database.Controller, hold constants.Hold, reason string, requester *database.Controller) error {
	for _, h := range controller.Holds {
		if h.Hold == hold {
			err := h.Delete()
			if err != nil {
				return err
			}
			err = LogAction(controller, fmt.Sprintf("Hold %s released: %s", hold, reason), requester)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New(fmt.Sprintf("controller %d does not have hold %s", controller.Id, hold))
}
