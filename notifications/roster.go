package notifications

import (
	"vatusa-api-v3/constants"
	"vatusa-api-v3/database"
)

func HomeControllerRemoved(controller *database.Controller, facility constants.Facility, actorName string, reason string) error {
	// TODO
	return nil
}
