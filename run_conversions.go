package main

import (
	"vatusa-api-v3/conversion"
	"vatusa-api-v3/conversion/legacydb"
	"vatusa-api-v3/database"
)

func main() {
	err := database.Connect()
	if err != nil {
		print(err)
		return
	}
	err = database.MigrateDB()
	if err != nil {
		print(err)
		return
	}
	err = legacydb.Connect()
	if err != nil {
		print(err)
		return
	}
	err = conversion.ConvertControllers()
	if err != nil {
		print(err)
		return
	}
	err = conversion.ConvertVisits()
	if err != nil {
		print(err)
		return
	}
}
