package database

import "gorm.io/gorm"

type WorkspaceAccountGroup struct {
	gorm.Model
	WorkspaceAccountID uint
	WorkspaceAccount   *WorkspaceAccount
	WorkspaceGroupID   uint
	WorkspaceGroup     *WorkspaceGroup
	IsManual           bool
}
