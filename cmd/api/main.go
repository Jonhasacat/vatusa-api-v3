package main

import (
	"github.com/VATUSA/api-v3/internal/database"
	v3 "github.com/VATUSA/api-v3/internal/v3/api"
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
	e := v3.App()
	e.Logger.Fatal(e.Start(":9001"))
}
