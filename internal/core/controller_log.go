package core

import (
	db "github.com/VATUSA/api-v3/internal/database"
	"github.com/VATUSA/api-v3/pkg/constants"
)

func LogAction(controller *db.Controller, message string, requester *db.Controller) error {
	// TODO
	return nil
}

func LogMessage(controller *db.Controller, visibility constants.LogVisibility, message string, requester db.Controller) error {
	// TODO
	return nil

}
