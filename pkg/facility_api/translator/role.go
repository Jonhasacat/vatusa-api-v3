package translator

import (
	db "github.com/VATUSA/api-v3/internal/database"
	"github.com/VATUSA/api-v3/pkg/facility_api/model"
)

func TranslateRole(role db.ControllerRole) model.Role {
	return model.Role{
		Facility: role.Facility,
		Role:     role.Role,
	}
}

func TranslateRoles(roles []db.ControllerRole) []model.Role {
	out := make([]model.Role, 0)
	for _, role := range roles {
		out = append(out, TranslateRole(role))
	}
	return out
}
