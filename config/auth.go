package config

import (
	"os"
)

var UserCookieName = os.Getenv("USER_COOKIE")
var JWTSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
