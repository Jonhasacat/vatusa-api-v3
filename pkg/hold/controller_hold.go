package hold

import (
	"errors"
	"fmt"
	"github.com/VATUSA/api-v3/pkg/action_log"
	"github.com/VATUSA/api-v3/pkg/database"
	"time"
)

func AddHold(controller *database.Controller, hold Hold, reason string, requester *database.Controller) error {
	record := database.ControllerHold{
		Controller: controller,
		Hold:       hold,
		ExpiresAt:  time.Time{},
	}
	err := record.Save()
	if err != nil {
		return err
	}
	err = action_log.LogAction(controller, fmt.Sprintf("Hold %s applied: %s", hold, reason), requester)
	return nil
}

func RemoveHold(controller *database.Controller, hold Hold, reason string, requester *database.Controller) error {
	for _, h := range controller.Holds {
		if h.Hold == hold {
			err := h.Delete()
			if err != nil {
				return err
			}
			err = action_log.LogAction(controller, fmt.Sprintf("Hold %s released: %s", hold, reason), requester)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New(fmt.Sprintf("controller %d does not have hold %s", controller.Id, hold))
}
