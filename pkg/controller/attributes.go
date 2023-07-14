package controller

import (
	"fmt"
	db "github.com/VATUSA/api-v3/pkg/database"
)

func CertificateName(c *db.Controller) string {
	if c.Certificate == nil {
		return "Unknown"
	}
	return fmt.Sprintf("%s %s", c.Certificate.FirstName, c.Certificate.LastName)
}

func DisplayName(c *db.Controller) string {
	return CertificateName(c)
}
