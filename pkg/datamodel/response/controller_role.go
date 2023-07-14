package response

import (
	"github.com/VATUSA/api-v3/pkg/database"
)

type ControllerRole struct {
	Facility string
	Role     string
}

func MakeControllerRoleResponse(r *database.ControllerRole) *ControllerRole {
	return &ControllerRole{
		Facility: string(r.Facility),
		Role:     string(r.Role),
	}
}
