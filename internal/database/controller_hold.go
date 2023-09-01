package database

import (
	"gorm.io/gorm"
	"time"
)

type ControllerHold struct {
	gorm.Model
	ControllerID uint64
	Controller   *Controller
	Hold         string
	ExpiresAt    *time.Time
}

func (h *ControllerHold) Save() error {
	result := DB.Save(h)
	if result.Error != nil {
		return result.Error
	}
	h.Controller.HookControllerUpdate()
	return nil
}

func (h *ControllerHold) Delete() error {
	controller := h.Controller
	result := DB.Delete(h)
	if result.Error != nil {
		return result.Error
	}
	controller.HookControllerUpdate()
	return nil
}
