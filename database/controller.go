package database

import (
	"fmt"
	"gorm.io/gorm"
	"time"
	"vatusa-api-v3/constants"
)

type Controller struct {
	Id                        uint64 `gorm:"primarykey;autoIncrement:false"`
	CertificateId             uint64
	Certificate               *Certificate       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Facility                  constants.Facility `gorm:"size:4"`
	FacilityJoin              *time.Time
	ATCRating                 constants.Rating
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

func FetchControllersByHomeFacility(facility constants.Facility) ([]Controller, error) {
	var controllers []Controller
	result := controllerQuery().
		Where("facility = ?", facility).
		Find(&controllers)
	if result.Error != nil {
		return nil, result.Error
	}
	return controllers, nil
}

func FetchControllersByVisitingFacility(facility constants.Facility) ([]Controller, error) {
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

func (c *Controller) CertificateName() string {
	if c.Certificate == nil {
		return "Unknown"
	}
	return fmt.Sprintf("%s %s", c.Certificate.FirstName, c.Certificate.LastName)
}

func (c *Controller) DisplayName() string {
	return c.CertificateName()
}

func (c *Controller) IsVisiting(fac constants.Facility) bool {
	for _, v := range c.Visits {
		if v.Facility == fac {
			return true
		}
	}
	return false
}

func (c *Controller) IsTransferEligible() bool {
	if !c.IsInDivision {
		return false
	}
	if !c.IsActive {
		return false
	}
	for _, h := range c.Holds {
		if h.ExpiresAt.Before(time.Now()) {
			DB.Delete(&h)
		} else {
			hold := constants.GetHold(h.Hold)
			if hold.PreventTransfer {
				return false
			}
		}
	}
	return true
}

func (c *Controller) IsVisitEligible() bool {
	if !c.IsInDivision && !c.IsApprovedExternalVisitor {
		return false
	}
	if !c.IsActive {
		return false
	}
	for _, h := range c.Holds {
		if h.ExpiresAt.Before(time.Now()) {
			DB.Delete(&h)
		} else {
			hold := constants.GetHold(h.Hold)
			if hold.PreventVisit {
				return false
			}
		}
	}
	return true
}

func (c *Controller) IsPromotionEligible() bool {
	if !c.IsInDivision {
		return false
	}
	if !c.IsActive {
		return false
	}
	for _, h := range c.Holds {
		if h.ExpiresAt.Before(time.Now()) {
			DB.Delete(&h)
		} else {
			hold := constants.GetHold(h.Hold)
			if hold.PreventPromotion {
				return false
			}
		}
	}
	// TODO
	return true
}

func (c *Controller) IsHomeController(facility constants.Facility) bool {
	return facility == c.Facility
}

func (c *Controller) IsVisitor(facility constants.Facility) bool {
	for _, v := range c.Visits {
		if v.Facility == facility {
			return true
		}
	}
	return false
}

func (c *Controller) IsOnFacilityRoster(facility constants.Facility) bool {
	return c.IsHomeController(facility) || c.IsVisitor(facility)
}

func (c *Controller) HasRole(role constants.Role, facility constants.Facility) bool {
	for _, r := range c.Roles {
		if r.Role == role && r.Facility == facility {
			return true
		}
	}
	return false
}

func (c *Controller) IsStaff(facility constants.Facility) bool {
	return c.IsSeniorStaff(facility) ||
		c.HasRole(constants.RoleWebMaster, facility) ||
		c.HasRole(constants.RoleFacilityEngineer, facility) ||
		c.HasRole(constants.RoleEventCoordinator, facility) ||
		c.IsTrainingStaff(facility)
}

func (c *Controller) IsSeniorStaff(facility constants.Facility) bool {
	return c.IsDivisionStaff() ||
		c.HasRole(constants.RoleAirTrafficManager, facility) ||
		c.HasRole(constants.RoleDeputyAirTrafficManager, facility) ||
		c.HasRole(constants.RoleTrainingAdministrator, facility)
}

func (c *Controller) IsATMOrDATM(facility constants.Facility) bool {
	return c.IsDivisionStaff() ||
		c.HasRole(constants.RoleAirTrafficManager, facility) ||
		c.HasRole(constants.RoleDeputyAirTrafficManager, facility)
}

func (c *Controller) IsDivisionStaff() bool {
	return c.HasRole(constants.RoleDivisionStaff, constants.Headquarters)
}

func (c *Controller) IsInstructor(facility constants.Facility) bool {
	return c.IsDivisionStaff() ||
		c.HasRole(constants.RoleInstructor, facility)
}

func (c *Controller) IsTrainingStaff(facility constants.Facility) bool {
	return c.IsSeniorStaff(facility) ||
		c.HasRole(constants.RoleInstructor, facility) ||
		c.HasRole(constants.RoleMentor, facility)
}
