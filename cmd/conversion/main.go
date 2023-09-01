package main

import (
	"github.com/VATUSA/api-v3/internal/conversion"
	"github.com/VATUSA/api-v3/internal/conversion/legacydb"
	"github.com/VATUSA/api-v3/internal/database"
)

func main() {
	err := database.Connect()
	if err != nil {
		println(err)
		return
	}
	err = database.MigrateDB()
	if err != nil {
		println(err)
		return
	}
	err = legacydb.Connect()
	if err != nil {
		println(err)
		return
	}
	err = conversion.ConvertControllers()
	if err != nil {
		println(err)
		return
	}
	err = conversion.ConvertVisits()
	if err != nil {
		println(err)
		return
	}
	err = conversion.ConvertRoles()
	if err != nil {
		println(err)
		return
	}
	err = conversion.ConvertPromotions()
	if err != nil {
		println(err)
		return
	}
}
