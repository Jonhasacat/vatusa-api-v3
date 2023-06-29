package database

import (
	"fmt"
	"gorm.io/gorm"
	"time"
	"vatusa-api-v3/constants"
)

type AcademyUser struct {
	gorm.Model
	ControllerID uint64
	Controller   *Controller
	MoodleUserID uint64
	LastSync     time.Time
}

func (a *AcademyUser) Sync() error {
	err := a.SyncUserInfo()
	if err != nil {
		return err
	}
	err = a.SyncCohorts()
	if err != nil {
		return err
	}
	return nil
}

func (a *AcademyUser) SyncUserInfo() error {
	// TODO
	return nil
}

func (a *AcademyUser) SyncRoles() error {
	err := a.clearRoles()
	if err != nil {
		return err
	}

	err = a.addFacilityRole(constants.MoodleRoleStudent, a.Controller.Facility)
	if err != nil {
		return err
	}

	if a.Controller.IsSeniorStaff(a.Controller.Facility) {
		err = a.addFacilityRole(constants.MoodleRoleInstructor, "USA")
		if err != nil {
			return err
		}
		err = a.addFacilityRole(constants.MoodleRoleFacilityAdmin, a.Controller.Facility)
		if err != nil {
			return err
		}
	}

	if a.Controller.HasRole(constants.RoleInstructor, a.Controller.Facility) {
		err = a.addFacilityRole(constants.MoodleRoleInstructor, "USA")
		if err != nil {
			return err
		}
		err = a.addFacilityRole(constants.MoodleRoleInstructor, a.Controller.Facility)
		if err != nil {
			return err
		}
	}

	for _, v := range a.Controller.Visits {
		err = a.addFacilityRole(constants.MoodleRoleStudent, v.Facility)
		if err != nil {
			return err
		}

		if a.Controller.HasRole(constants.RoleInstructor, v.Facility) {
			err = a.addFacilityRole(constants.MoodleRoleInstructor, v.Facility)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (a *AcademyUser) clearRoles() error {
	// TODO
	return nil
}

func (a *AcademyUser) addFacilityRole(roleID uint64, facility constants.Facility) error {
	return a.addCategoryRole(roleID, string(facility))
}

func (a *AcademyUser) addCategoryRole(roleID uint64, category string) error {
	// TODO
	return nil
}

func (a *AcademyUser) addCourseRole(roleID uint64, courseID uint64) error {
	// TODO
	return nil
}

func (a *AcademyUser) SyncCohorts() error {
	var cohorts []string

	if !a.Controller.IsActive || a.Controller.Facility == constants.Inactive {
		return nil
	}

	// Home Facility Cohort
	cohorts = append(cohorts, string(a.Controller.Facility))

	if a.Controller.HasRole("MTR", a.Controller.Facility) {
		cohorts = append(cohorts, "TNG")
		cohorts = append(cohorts, "MTR")
		cohorts = append(cohorts, fmt.Sprintf("%s-MTR", a.Controller.Facility))
	}
	if a.Controller.HasRole("INS", a.Controller.Facility) {
		cohorts = append(cohorts, "TNG")
		cohorts = append(cohorts, "INS")
		cohorts = append(cohorts, fmt.Sprintf("%s-INS", a.Controller.Facility))
	}

	// Visitor Cohorts
	for _, v := range a.Controller.Visits {
		cohorts = append(cohorts, fmt.Sprintf("%s-V", v.Facility))
		if a.Controller.HasRole("MTR", v.Facility) {
			cohorts = append(cohorts, fmt.Sprintf("%s-MTR", v.Facility))
		}
		if a.Controller.HasRole("INS", v.Facility) {
			cohorts = append(cohorts, fmt.Sprintf("%s-INS", v.Facility))
		}
	}

	// Rating Cohort
	if a.Controller.ATCRating > constants.RatingC1 {
		cohorts = append(cohorts, "C1+")
		cohorts = append(cohorts, fmt.Sprintf("%s-C1+", a.Controller.Facility))
	} else {
		cohorts = append(cohorts, constants.RatingShortMap[a.Controller.ATCRating])
		cohorts = append(cohorts, fmt.Sprintf(
			"%s-%s", a.Controller.Facility, constants.RatingShortMap[a.Controller.ATCRating]))
	}

	return a.setUserCohorts(cohorts)
}

func (a *AcademyUser) setUserCohorts(cohorts []string) error {
	err := a.clearCohorts()
	if err != nil {
		return err
	}

	for _, c := range cohorts {
		err = a.addCohort(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *AcademyUser) clearCohorts() error {
	// TODO
	return nil
}

func (a *AcademyUser) addCohort(cohort string) error {
	// TODO
	return nil
}
