package model

import "vatusa-api-v3/database"

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
