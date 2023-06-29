package main

import (
	"vatusa-api-v3/database"
	v3 "vatusa-api-v3/v3/api"
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
