package operations

import (
	"vatusa-api-v3/constants"
	db "vatusa-api-v3/database"
)

func LogAction(controller *db.Controller, message string, requester *db.Controller) error {
	// TODO
	return nil
}

func LogMessage(controller *db.Controller, visibility constants.LogVisibility, message string, requester db.Controller) error {
	// TODO
	return nil

}
