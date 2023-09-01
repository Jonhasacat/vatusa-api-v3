package database

import (
	"errors"
	"fmt"
	"github.com/VATUSA/api-v3/pkg/constants"
	"gorm.io/gorm"
)

type ControllerRosterRequest struct {
	gorm.Model
	ControllerID uint64
	Controller   *Controller
	Facility     constants.Facility
	RequestType  string // constants.RequestTransfer or constants.RequestVisit
	Reason       string
	RequesterID  uint64
	Requester    *Controller
	Status       string
	StatusReason *string
	AdminID      uint64
	Admin        *Controller
}

func requestsQuery() *gorm.DB {
	return DB.
		Model(&ControllerRosterRequest{}).
		Joins("Controller").
		Preload("Controller.Holds").
		Preload("Controller.RatingChanges").
		Preload("Controller.Roles").
		Preload("Controller.Visits")
}

func FetchRequestById(id uint) (*ControllerRosterRequest, error) {
	var request *ControllerRosterRequest
	requestsQuery().First(request, id)
	if request == nil {
		return nil, errors.New(fmt.Sprintf("Roster Request %d does not exist", id))
	}
	return request, nil
}

func FetchPendingRequestsByFacility(facility string) ([]ControllerRosterRequest, error) {
	var requests []ControllerRosterRequest
	result := requestsQuery().
		Where("facility = ?", facility).
		Find(&requests)
	if result.Error != nil {
		return nil, result.Error
	}
	return requests, nil
}
