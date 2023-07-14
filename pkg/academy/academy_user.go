package academy

import (
	"fmt"
	"github.com/VATUSA/api-v3/pkg/controller"
	db "github.com/VATUSA/api-v3/pkg/database"
	"github.com/VATUSA/api-v3/pkg/facility"
	"github.com/VATUSA/api-v3/pkg/rating"
	"github.com/VATUSA/api-v3/pkg/role"
)

func CreateAcademyUser(au *db.AcademyUser) error {
	// TODO
	return nil
}

func SyncAcademyUser(au *db.AcademyUser) error {
	// TODO
	return nil
}

func SyncRoles(a *db.AcademyUser) error {
	err := clearRoles(a)
	if err != nil {
		return err
	}

	err = addFacilityRole(a, RoleStudent, a.Controller.Facility)
	if err != nil {
		return err
	}

	if controller.IsSeniorStaff(a.Controller, a.Controller.Facility) {
		err = addFacilityRole(a, RoleInstructor, "USA")
		if err != nil {
			return err
		}
		err = addFacilityRole(a, RoleFacilityAdmin, a.Controller.Facility)
		if err != nil {
			return err
		}
	}

	if controller.HasRole(a.Controller, role.RoleInstructor, a.Controller.Facility) {
		err = addFacilityRole(a, RoleInstructor, "USA")
		if err != nil {
			return err
		}
		err = addFacilityRole(a, RoleInstructor, a.Controller.Facility)
		if err != nil {
			return err
		}
	}

	for _, v := range a.Controller.Visits {
		err = addFacilityRole(a, RoleStudent, v.Facility)
		if err != nil {
			return err
		}

		if controller.HasRole(a.Controller, role.RoleInstructor, v.Facility) {
			err = addFacilityRole(a, RoleInstructor, v.Facility)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func clearRoles(a *db.AcademyUser) error {
	// TODO
	return nil
}

func addFacilityRole(a *db.AcademyUser, roleID uint64, facility facility.Facility) error {
	return addCategoryRole(a, roleID, string(facility))
}

func addCategoryRole(a *db.AcademyUser, roleID uint64, category string) error {
	// TODO
	return nil
}

func addCourseRole(a *db.AcademyUser, roleID uint64, courseID uint64) error {
	// TODO
	return nil
}

func SyncCohorts(a *db.AcademyUser) error {
	var cohorts []string

	if !a.Controller.IsActive || a.Controller.Facility == facility.Inactive {
		return nil
	}

	// Home Facility Cohort
	cohorts = append(cohorts, string(a.Controller.Facility))

	if controller.HasRole(a.Controller, "MTR", a.Controller.Facility) {
		cohorts = append(cohorts, "TNG")
		cohorts = append(cohorts, "MTR")
		cohorts = append(cohorts, fmt.Sprintf("%s-MTR", a.Controller.Facility))
	}
	if controller.HasRole(a.Controller, "INS", a.Controller.Facility) {
		cohorts = append(cohorts, "TNG")
		cohorts = append(cohorts, "INS")
		cohorts = append(cohorts, fmt.Sprintf("%s-INS", a.Controller.Facility))
	}

	// Visitor Cohorts
	for _, v := range a.Controller.Visits {
		cohorts = append(cohorts, fmt.Sprintf("%s-V", v.Facility))
		if controller.HasRole(a.Controller, "MTR", v.Facility) {
			cohorts = append(cohorts, fmt.Sprintf("%s-MTR", v.Facility))
		}
		if controller.HasRole(a.Controller, "INS", v.Facility) {
			cohorts = append(cohorts, fmt.Sprintf("%s-INS", v.Facility))
		}
	}

	// Rating Cohort
	if rating.Rating(a.Controller.ATCRating) > rating.C1 {
		cohorts = append(cohorts, "C1+")
		cohorts = append(cohorts, fmt.Sprintf("%s-C1+", a.Controller.Facility))
	} else {
		cohorts = append(cohorts, rating.ShortFromInt(a.Controller.ATCRating))
		cohorts = append(cohorts, fmt.Sprintf(
			"%s-%s", a.Controller.Facility, rating.ShortFromInt(a.Controller.ATCRating)))
	}

	return setUserCohorts(a, cohorts)
}

func setUserCohorts(a *db.AcademyUser, cohorts []string) error {
	err := clearCohorts(a)
	if err != nil {
		return err
	}

	for _, c := range cohorts {
		err = addCohort(a, c)
		if err != nil {
			return err
		}
	}
	return nil
}

func clearCohorts(a *db.AcademyUser) error {
	// TODO
	return nil
}

func addCohort(a *db.AcademyUser, cohort string) error {
	// TODO
	return nil
}
