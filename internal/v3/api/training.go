package api

import (
	"fmt"
	database2 "github.com/VATUSA/api-v3/pkg/database"
	"github.com/VATUSA/api-v3/pkg/datamodel/response"
	"github.com/VATUSA/api-v3/pkg/facility"
	"github.com/VATUSA/api-v3/pkg/roster"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type TrainingRecordRequest struct {
	StudentCID        uint64
	InstructorCID     uint64
	SessionTime       time.Time
	Facility          facility.Facility
	Position          string
	DurationMinutes   uint64
	AircraftMovements uint64
	Score             uint64
	InstructionMethod string
	// Values: constants.InstructionMethodLecture, constants.InstructionMethodSweatbox, constants.InstructionMethodLive
	IsOTSRecommendation bool
	IsSoloCertGranted   bool
	SoloCertPosition    *string // Required if IsSoloCertGranted == true

}

func CreateTrainingRecord(c echo.Context) error {
	var request TrainingRecordRequest
	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if request.IsSoloCertGranted && request.SoloCertPosition == nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			"SoloCertPosition must be specified when IsSoloCertGranted = true")
	}
	student, err := database2.FetchControllerByCID(request.StudentCID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if !roster.IsOnFacilityRoster(student, request.Facility) {
		return echo.NewHTTPError(http.StatusBadRequest,
			fmt.Sprintf("Controller %d is not on the %s roster", request.StudentCID, request.Facility))
	}
	instructor, err := database2.FetchControllerByCID(request.InstructorCID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if !roster.IsOnFacilityRoster(instructor, request.Facility) {
		return echo.NewHTTPError(http.StatusBadRequest,
			fmt.Sprintf("Controller %d is not on the %s roster", request.InstructorCID, request.Facility))
	}
	record := database2.TrainingRecord{
		StudentCID:          request.StudentCID,
		Student:             student,
		InstructorCID:       request.InstructorCID,
		Instructor:          instructor,
		SessionTime:         request.SessionTime,
		Facility:            request.Facility,
		Position:            request.Position,
		DurationMinutes:     request.DurationMinutes,
		AircraftMovements:   request.AircraftMovements,
		Score:               request.Score,
		InstructionMethod:   request.InstructionMethod,
		IsOTSRecommendation: request.IsOTSRecommendation,
		IsSoloCertGranted:   request.IsSoloCertGranted,
		SoloCertPosition:    request.SoloCertPosition,
	}
	record.Save()
	return c.JSON(http.StatusOK, response.MakeTrainingRecordResponse(&record))
}

func ModifyTrainingRecord(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var request TrainingRecordRequest
	err = c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if request.IsSoloCertGranted && request.SoloCertPosition == nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			"SoloCertPosition must be specified when IsSoloCertGranted = true")
	}
	student, err := database2.FetchControllerByCID(request.StudentCID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if !roster.IsOnFacilityRoster(student, request.Facility) {
		return echo.NewHTTPError(http.StatusBadRequest,
			fmt.Sprintf("Controller %d is not on the %s roster", request.StudentCID, request.Facility))
	}
	instructor, err := database2.FetchControllerByCID(request.InstructorCID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if !roster.IsOnFacilityRoster(instructor, request.Facility) {
		return echo.NewHTTPError(http.StatusBadRequest,
			fmt.Sprintf("Controller %d is not on the %s roster", request.InstructorCID, request.Facility))
	}
	record, err := database2.FetchTrainingRecordByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	record.StudentCID = request.StudentCID
	record.Student = student
	record.InstructorCID = request.InstructorCID
	record.Instructor = instructor
	record.SessionTime = request.SessionTime
	record.Facility = request.Facility
	record.Position = request.Position
	record.DurationMinutes = request.DurationMinutes
	record.AircraftMovements = request.AircraftMovements
	record.Score = request.Score
	record.InstructionMethod = request.InstructionMethod
	record.IsOTSRecommendation = request.IsOTSRecommendation
	record.IsSoloCertGranted = request.IsSoloCertGranted
	record.SoloCertPosition = request.SoloCertPosition

	record.Save()
	return c.JSON(http.StatusOK, response.MakeTrainingRecordResponse(record))
}

func DeleteTrainingRecord(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	record, err := database2.FetchTrainingRecordByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	record.Delete()
	return nil
}

func GetControllerTrainingRecords(c echo.Context) error {
	cidParam := c.Param("cid")
	cid, err := strconv.ParseUint(cidParam, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	records, err := database2.FetchTrainingRecordsByCID(cid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var output []*response.TrainingRecord
	for _, v := range records {
		output = append(output, response.MakeTrainingRecordResponse(&v))
	}
	return c.JSON(http.StatusOK, output)
}

func GetFacilityTrainingRecords(c echo.Context) error {
	facility := c.Param("facility")
	records, err := database2.FetchTrainingRecordsByFacility(facility)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var output []*response.TrainingRecord
	for _, v := range records {
		output = append(output, response.MakeTrainingRecordResponse(&v))
	}
	return c.JSON(http.StatusOK, output)
}
