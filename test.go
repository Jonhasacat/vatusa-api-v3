package main

import (
	"github.com/VATUSA/api-v3/internal/conversion/legacydb"
	"github.com/VATUSA/api-v3/internal/tasks"
	db "github.com/VATUSA/api-v3/pkg/database"
)

func main() {
	err := db.Connect()
	if err != nil {
		return
	}
	err = db.MigrateDB()
	if err != nil {
		return
	}
	err = legacydb.Connect()
	if err != nil {
		return
	}
	tasks.SyncRosterFromVATSIM()
}
