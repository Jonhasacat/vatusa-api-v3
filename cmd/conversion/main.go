package main

import (
	"github.com/VATUSA/api-v3/internal/conversion"
	"github.com/VATUSA/api-v3/internal/conversion/legacydb"
	db "github.com/VATUSA/api-v3/pkg/database"
)

func main() {
	err := db.Connect()
	if err != nil {
		println(err)
		return
	}
	err = db.MigrateDB()
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
