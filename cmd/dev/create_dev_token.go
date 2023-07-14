package main

import (
	"fmt"
	"github.com/VATUSA/api-v3/internal/conversion/legacydb"
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
	err = legacydb.Connect()
	if err != nil {
		return
	}
	user, err := database2.CreateAPIUser("Development API User", "*")
	if err != nil {
		return
	}
	token, err := database2.GenerateAPIToken(user, nil)
	if err != nil {
		return
	}
	print(fmt.Sprintf("Token Generated: %s", token.Token))
}
