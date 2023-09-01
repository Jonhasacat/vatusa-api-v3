package database

import (
	"errors"
	"fmt"
	"github.com/VATUSA/api-v3/pkg/constants"
	"gorm.io/gorm"
	"time"
)

type TrainingRecord struct {
	gorm.Model
	StudentCID          uint64
	Student             *Controller
	InstructorCID       uint64
	Instructor          *Controller
	SessionTime         time.Time
	Facility            constants.Facility
	Position            string
	DurationMinutes     uint64
	AircraftMovements   uint64
	Score               uint64
	InstructionMethod   string
	IsOTSRecommendation bool
	IsSoloCertGranted   bool
	SoloCertPosition    *string
}

func trainingRecordQuery() *gorm.DB {
	return DB.
		Model(&TrainingRecord{}).
		Joins("Student").
		Joins("Instructor")
}

func FetchTrainingRecordByID(id uint64) (*TrainingRecord, error) {
	var record *TrainingRecord
	trainingRecordQuery().First(&record, id)
	if record == nil {
		return nil, errors.New(fmt.Sprintf("Training Record %d does not exist", id))
	}
	return record, nil
}

func FetchTrainingRecordsByCID(cid uint64) ([]TrainingRecord, error) {
	var records []TrainingRecord
	result := trainingRecordQuery().
		Where("StudentCID = ?", cid).
		Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}
	return records, nil
}

func FetchTrainingRecordsByFacility(facility string) ([]TrainingRecord, error) {
	var records []TrainingRecord
	result := trainingRecordQuery().
		Where("Facility = ?", facility).
		Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}
	return records, nil
}

func (tr *TrainingRecord) Save() {
	DB.Save(tr)
}

func (tr *TrainingRecord) Delete() {
	DB.Delete(tr)
}
