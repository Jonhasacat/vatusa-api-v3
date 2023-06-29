package model

import (
	"time"
	"vatusa-api-v3/database"
)

type TrainingRecord struct {
	ID                  uint
	StudentCID          uint64
	Student             *Controller
	InstructorCID       uint64
	Instructor          *Controller
	SessionTime         time.Time
	Facility            string
	Position            string
	DurationMinutes     uint64
	AircraftMovements   uint64
	Score               uint64
	InstructionMethod   string
	IsOTSRecommendation bool
	IsSoloCertGranted   bool
	SoloCertPosition    *string
}

func MakeTrainingRecordResponse(tr *database.TrainingRecord) *TrainingRecord {
	record := &TrainingRecord{
		ID:                  tr.ID,
		StudentCID:          tr.StudentCID,
		Student:             MakeController(tr.Student),
		InstructorCID:       tr.InstructorCID,
		Instructor:          MakeController(tr.Instructor),
		SessionTime:         tr.SessionTime,
		Facility:            tr.Facility,
		Position:            tr.Position,
		DurationMinutes:     tr.DurationMinutes,
		AircraftMovements:   tr.AircraftMovements,
		Score:               tr.Score,
		InstructionMethod:   tr.InstructionMethod,
		IsOTSRecommendation: tr.IsOTSRecommendation,
		IsSoloCertGranted:   tr.IsSoloCertGranted,
		SoloCertPosition:    tr.SoloCertPosition,
	}
	return record
}
