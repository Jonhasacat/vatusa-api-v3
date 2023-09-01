package main

import (
	"fmt"
	"github.com/VATUSA/api-v3/internal/conversion/legacydb"
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
