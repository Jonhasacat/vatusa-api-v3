package main

import (
	"fmt"
	"vatusa-api-v3/conversion/legacydb"
	"vatusa-api-v3/database"
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
	user, err := database.CreateAPIUser("Development API User", "*")
	if err != nil {
		return
	}
	token, err := database.GenerateAPIToken(user, nil)
	if err != nil {
		return
	}
	print(fmt.Sprintf("Token Generated: %s", token.Token))
}
