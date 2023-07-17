package conversion

import (
	"fmt"
	"github.com/VATUSA/api-v3/internal/conversion/legacydb"
	"github.com/VATUSA/api-v3/pkg/async"
	db "github.com/VATUSA/api-v3/pkg/database"
	"github.com/VATUSA/api-v3/pkg/facility"
	"github.com/VATUSA/api-v3/pkg/rating"
	"github.com/VATUSA/api-v3/pkg/role"
	"time"
)

const (
	WorkerCount    = 10
	RecordsPerPage = 100
)

func ProcessLegacyController(legacy legacydb.Controller) error {
	certificate, _ := db.FetchCertificateByID(legacy.CID)
	if certificate == nil {
		certificate = &db.Certificate{
			ID:                     legacy.CID,
			FirstName:              legacy.FName,
			LastName:               legacy.LName,
			Email:                  legacy.Email,
			Rating:                 legacy.Rating,
			PilotRating:            0,
			MilitaryRating:         0,
			SuspendDate:            nil,
			RegistrationDate:       nil,
			Region:                 nil,
			Division:               nil,
			SubDivision:            nil,
			LastRatingChange:       nil,
			CertificateUpdateStamp: time.Unix(0, 0),
		}
		err := certificate.Save()
		if err != nil {
			return err
		}
	}
	controller, _ := db.FetchControllerByCID(legacy.CID)
	if controller == nil {
		controller = &db.Controller{
			Id:                        legacy.CID,
			CertificateId:             legacy.CID,
			Certificate:               certificate,
			Facility:                  legacy.Facility,
			FacilityJoin:              legacy.FacilityJoin,
			ATCRating:                 legacy.Rating,
			LastPromotion:             nil,
			IsInDivision:              legacy.FlagHomeController,
			IsApprovedExternalVisitor: false,
			IsActive:                  false,
			DiscordId:                 legacy.DiscordId,
		}
		db.DB.Create(&controller)
	} else {

	}
	if controller.IsInDivision &&
		facility.IsRosterFacility(controller.Facility) &&
		(controller.ATCRating == rating.I1 || controller.ATCRating == rating.I3) {
		if !role.HasRole(controller, role.Instructor, controller.Facility) {
			err := role.AddRole(controller, role.Instructor, controller.Facility, nil)
			if err != nil {
				return err
			}
		}
	}
	if controller.ATCRating < rating.Observer || controller.ATCRating > rating.I3 {
		if controller.ATCRating > rating.I3 && role.HasRole(controller, role.TrainingAdministrator, controller.Facility) {
			controller.ATCRating = rating.I3
		} else if controller.ATCRating > rating.I3 && role.HasRole(controller, role.Instructor, controller.Facility) {
			controller.ATCRating = rating.I1
		}
		// TODO: Determine which ATCRating they should be
	}
	return nil
}

func LoadLegacyControllerPage(page int) ([]legacydb.Controller, error) {
	println(fmt.Sprintf("loading page %d", page))
	var controllers []legacydb.Controller
	result := legacydb.DB.Model(legacydb.Controller{}).Limit(RecordsPerPage).Offset(RecordsPerPage * page).Find(&controllers)
	if result.Error != nil {
		return nil, result.Error
	}
	return controllers, nil
}

func ConvertControllerWorker(offset int) {
	for i := 0; true; i++ {
		page := (i * WorkerCount) + offset
		controllers, err := LoadLegacyControllerPage(page)
		if err != nil {
			println(fmt.Sprintf("Error in worker %d: %s", offset, err))
		}
		if len(controllers) == 0 {
			println(fmt.Sprintf("Worker %d fetched no controllers on page %d", offset, page))
			break
		}
		for _, controller := range controllers {
			err = ProcessLegacyController(controller)
			if err != nil {
				println(fmt.Sprintf("Error while processing CID %d: %s", controller.CID, err.Error()))
			}
		}
	}
	println(fmt.Sprintf("Worker %d finished", offset))
}

func ConvertControllers() error {
	async.SpawnWorkers(WorkerCount, ConvertControllerWorker)
	return nil
}
