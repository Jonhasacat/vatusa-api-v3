package core

import (
	"github.com/VATUSA/api-v3/pkg/constants"
	db "github.com/VATUSA/api-v3/pkg/database"
)

func LogAction(controller *db.Controller, message string, requester *db.Controller) error {
	// TODO
	return nil
}

func LogMessage(controller *db.Controller, visibility constants.LogVisibility, message string, requester db.Controller) error {
	// TODO
	return nil

}
