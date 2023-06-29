package operations

import (
	"errors"
	"fmt"
	"time"
	"vatusa-api-v3/constants"
	db "vatusa-api-v3/database"
)

func AddHold(controller *db.Controller, hold constants.Hold, reason string, requester *db.Controller) error {
	record := db.ControllerHold{
		Controller: controller,
		Hold:       hold,
		ExpiresAt:  time.Time{},
	}
	err := record.Save()
	if err != nil {
		return err
	}
	err = LogAction(controller, fmt.Sprintf("Hold %s applied: %s", hold, reason), requester)
	return nil
}

func RemoveHold(controller *db.Controller, hold constants.Hold, reason string, requester *db.Controller) error {
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
