package main

import (
	"github.com/VATUSA/api-v3/internal/api/facility"
	"github.com/VATUSA/api-v3/internal/database"
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
	e := facility.App()
	e.Logger.Fatal(e.Start(":9002"))
}
