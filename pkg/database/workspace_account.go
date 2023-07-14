package database

import "gorm.io/gorm"

type WorkspaceAccount struct {
	gorm.Model
	ControllerID uint64
	Controller   *Controller
	Username     *string
	IsEnabled    bool
	CanSelfServe bool
	WorkspaceID  *string
	IsManual     bool
}
