package main

import (
	"github.com/VATUSA/api-v3/internal/conversion/legacydb"
	"github.com/VATUSA/api-v3/internal/tasks"
	database2 "github.com/VATUSA/api-v3/pkg/database"
)

func main() {
	err := database2.Connect()
	if err != nil {
		return
	}
	err = database2.MigrateDB()
	if err != nil {
		return
	}
	err = legacydb.Connect()
	if err != nil {
		return
	}
	page, err := tasks.FetchDivisionRosterPage(0)
	print(page)
}
