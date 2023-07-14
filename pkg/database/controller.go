package database

import (
	"github.com/VATUSA/api-v3/pkg/facility"
	"gorm.io/gorm"
	"time"
)

type Controller struct {
	Id                        uint64 `gorm:"primarykey;autoIncrement:false"`
	CertificateId             uint64
	Certificate               *Certificate      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Facility                  facility.Facility `gorm:"size:4"`
	FacilityJoin              *time.Time
	ATCRating                 int
	LastPromotion             *time.Time
	IsInDivision              bool
	IsApprovedExternalVisitor bool
	IsActive                  bool
	DiscordId                 *string `gorm:"size:40"`
	Holds                     []ControllerHold
	RatingChanges             []RatingChange
	Visits                    []ControllerVisit
	Transfers                 []Transfer
	Roles                     []ControllerRole
}

func controllerQuery() *gorm.DB {
	query := DB.
		Model(&Controller{}).
		Joins("Certificate").
		Preload("Holds").
		Preload("RatingChanges").
		Preload("Visits").
		Preload("Roles")
	return query
}

func FetchControllerByCID(cid uint64) (*Controller, error) {
	var model *Controller
	result := controllerQuery().Where(&Controller{Id: cid}).First(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}

func FetchControllersByHomeFacility(facility facility.Facility) ([]Controller, error) {
	var controllers []Controller
	result := controllerQuery().
		Where("facility = ?", facility).
		Find(&controllers)
	if result.Error != nil {
		return nil, result.Error
	}
	return controllers, nil
}

func FetchControllersByVisitingFacility(facility facility.Facility) ([]Controller, error) {
	var visits []ControllerVisit
	visitResult := DB.Model(&ControllerVisit{}).
		Where("facility = ?", facility).
		Find(&visits)
	if visitResult.Error != nil {
		return nil, visitResult.Error
	}
	var ids []uint64
	for _, v := range visits {
		ids = append(ids, v.ControllerID)
	}

	var controllers []Controller
	result := controllerQuery().Where(ids).Find(&controllers)
	if result.Error != nil {
		return nil, result.Error
	}
	return controllers, nil
}

func (c *Controller) HookControllerUpdate() {

}

func (c *Controller) Save() {
	DB.Save(c)
	c.HookControllerUpdate()
}
