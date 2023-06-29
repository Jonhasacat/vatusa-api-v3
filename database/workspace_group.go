package database

import "gorm.io/gorm"

type WorkspaceGroup struct {
	gorm.Model
	WorkspaceDomainID uint
	WorkspaceDomain   *WorkspaceDomain
	Group             string
	DisplayName       string
	Facility          string
	IsManual          bool
}
