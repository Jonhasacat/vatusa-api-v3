package response

import (
	"github.com/VATUSA/api-v3/internal/database"
)

type ControllerRole struct {
	Facility string
	Role     string
}

func MakeControllerRoleResponse(r *database.ControllerRole) *ControllerRole {
	return &ControllerRole{
		Facility: r.Facility,
		Role:     r.Role,
	}
}
