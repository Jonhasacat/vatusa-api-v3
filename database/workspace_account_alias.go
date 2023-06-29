package database

import "gorm.io/gorm"

type WorkspaceAccountAlias struct {
	gorm.Model
	WorkspaceAccountID uint
	WorkspaceAccount   *WorkspaceAccount
	WorkspaceDomainID  uint
	WorkspaceDomain    *WorkspaceDomain
	Alias              string
	IsManual           bool
}
