package main

import (
	v3 "github.com/VATUSA/api-v3/internal/v3/api"
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
	e := v3.App()
	e.Logger.Fatal(e.Start(":9001"))
}
