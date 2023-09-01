package main

import (
	"github.com/VATUSA/api-v3/internal/conversion/legacydb"
	"github.com/VATUSA/api-v3/internal/database"
	"github.com/VATUSA/api-v3/internal/tasks"
)

func main() {
	err := database.Connect()
	if err != nil {
		return
	}
	err = database.MigrateDB()
	if err != nil {
		return
	}
	err = legacydb.Connect()
	if err != nil {
		return
	}
	tasks.SyncRosterFromVATSIM()
}
