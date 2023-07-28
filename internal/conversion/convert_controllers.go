package conversion

import (
	"fmt"
	"github.com/VATUSA/api-v3/internal/conversion/legacydb"
	"github.com/VATUSA/api-v3/internal/core"
	"github.com/VATUSA/api-v3/pkg/async"
	"github.com/VATUSA/api-v3/pkg/constants"
	db "github.com/VATUSA/api-v3/pkg/database"
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
		constants.IsRosterFacility(controller.Facility) &&
		(controller.ATCRating == constants.I1 || controller.ATCRating == constants.I3) {
		if !core.HasRole(controller, constants.Instructor, controller.Facility) {
			err := core.AddRole(controller, constants.Instructor, controller.Facility, nil)
			if err != nil {
				return err
			}
		}
	}
	if controller.ATCRating < constants.Observer || controller.ATCRating > constants.I3 {
		if controller.ATCRating > constants.I3 && core.HasRole(controller, constants.TrainingAdministrator, controller.Facility) {
			controller.ATCRating = constants.I3
		} else if controller.ATCRating > constants.I3 && core.HasRole(controller, constants.Instructor, controller.Facility) {
			controller.ATCRating = constants.I1
		}
		legacyPromotions, err := LoadLegacyPromotionsByCID(controller.Id)
		if err != nil {
			return err
		}
		if len(legacyPromotions) == 0 {
			controller.ATCRating = constants.Observer
			// This assumption isn't guaranteed
			// it is possible that someone promoted prior to ~2008 will be set as OBS and need to be corrected
		} else {
			var lastPromotion *legacydb.Promotion
			// Has Legacy Promotions
			for _, p := range legacyPromotions {
				if lastPromotion == nil || p.CreatedAt.After(lastPromotion.CreatedAt) {
					lastPromotion = &p
				}
			}
			controller.ATCRating = lastPromotion.ToRating
		}
		if certificate.Rating == constants.Administrator && controller.ATCRating == constants.Observer {
			// Assume ADM are actually C1, if we don't have any other history
			controller.ATCRating = constants.C1
		}
		if certificate.Rating == constants.Inactive && controller.ATCRating > constants.C3 {
			err := core.ChangeRating(controller, constants.C1, nil)
			if err != nil {
				return err
			}
		}
		// TODO: Determine which ATCRating they should be
	}
	err := controller.Save()
	if err != nil {
		return err
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
