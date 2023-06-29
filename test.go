package main

import (
	"vatusa-api-v3/conversion/legacydb"
	"vatusa-api-v3/database"
	"vatusa-api-v3/tasks"
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
	page, err := tasks.FetchDivisionRosterPage(0)
	print(page)
}
