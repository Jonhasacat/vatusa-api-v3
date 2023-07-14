package main

import (
	conversion2 "github.com/VATUSA/api-v3/internal/conversion"
	"github.com/VATUSA/api-v3/internal/conversion/legacydb"
	database2 "github.com/VATUSA/api-v3/pkg/database"
)

func main() {
	err := database2.Connect()
	if err != nil {
		print(err)
		return
	}
	err = database2.MigrateDB()
	if err != nil {
		print(err)
		return
	}
	err = legacydb.Connect()
	if err != nil {
		print(err)
		return
	}
	err = conversion2.ConvertControllers()
	if err != nil {
		print(err)
		return
	}
	err = conversion2.ConvertVisits()
	if err != nil {
		print(err)
		return
	}
}
