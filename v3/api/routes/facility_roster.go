package routes

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"vatusa-api-v3/auth"
	"vatusa-api-v3/constants"
	"vatusa-api-v3/database"
	"vatusa-api-v3/operations"
	"vatusa-api-v3/v3/model"
)

// GetFacilityRoster godoc
//	@Summary	Get Facility Roster
//	@Tags		facility
//	@Param		facility	path	int	true	"Facility"
//	@Router		/v3/facility/{id}/roster [get]
//	@Success	200	{array}		model.ControllerLimited
//	@Success	200	{array}		model.Controller
//	@Failure	400	{object}	model.ErrorMessage
func GetFacilityRoster(c echo.Context) error {
	facility := constants.Facility(c.Param("facility"))
	controllers, err := database.FetchControllersByHomeFacility(facility)
	visitors, err := database.FetchControllersByVisitingFacility(facility)
	controllers = append(controllers, visitors...)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to fetch data")
	}
	if auth.CanReadControllerSensitiveData(c) {
		var output []*model.Controller
		for _, v := range controllers {
			output = append(output, model.MakeController(&v))
		}
		return c.JSON(http.StatusOK, output)
	} else if auth.IsAuthenticated(c) {
		var output []*model.Controller
		for _, v := range controllers {
			output = append(output, model.MakeControllerRedacted(&v))
		}
		return c.JSON(http.StatusOK, output)
	} else {
		var output []*model.ControllerLimited
		for _, v := range controllers {
			output = append(output, model.MakeControllerLimited(&v))
		}
		return c.JSON(http.StatusOK, output)
	}
}

type RemoveFromRosterRequest struct {
	CID          uint64  `json:"cid"`
	Reason       string  `json:"reason"`
	RequesterCID *uint64 `json:"requester_cid"`
}

func RemoveFromRoster(c echo.Context) error {
	facility := constants.Facility(c.Param("facility"))
	var request RemoveFromRosterRequest
	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Payload is incorrectly formatted")
	}
	controller, err := database.FetchControllerByCID(request.CID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("CID %d does not exist", request.CID))
	}
	var requester *database.Controller
	if request.RequesterCID != nil {
		requester, err = database.FetchControllerByCID(*request.RequesterCID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("CID %d does not exist", request.RequesterCID))
		}
	}
	if controller.Facility == facility {
		// Home Removal
		err := operations.RemoveFromFacility(controller, requester, request.Reason)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	} else if controller.IsVisiting(facility) {
		err := operations.RemoveVisitor(controller, facility, requester, request.Reason)
		if err != nil {

		}
	} else {
		return echo.NewHTTPError(
			http.StatusBadRequest, fmt.Sprintf("Controller %d is not on the %s roster", request.CID, facility))
	}
	return nil
}

func GetPendingRosterRequests(c echo.Context) error {
	facility := c.Param("facility")
	requests, err := database.FetchPendingRequestsByFacility(facility)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	var output []*model.ControllerRosterRequest
	for _, v := range requests {
		output = append(output, model.MakeControllerRosterRequestResponse(&v))
	}
	return c.JSON(http.StatusOK, output)
}

type ProcessRosterRequestRequest struct {
	ID           uint
	Accept       bool
	Reason       *string
	RequesterCID uint64
}

func ProcessRosterRequest(c echo.Context) error {
	facility := constants.Facility(c.Param("facility"))
	var request ProcessRosterRequestRequest
	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Payload is incorrectly formatted")
	}
	record, err := database.FetchRequestById(request.ID)
	if record.Status != constants.RequestPending {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Roster Request %d is not Pending", request.ID))
	}
	if record.Facility != facility {
		return echo.NewHTTPError(
			http.StatusBadRequest, fmt.Sprintf("Roster Request %d is not for facility %s", request.ID, facility))
	}
	requester, err := database.FetchControllerByCID(request.RequesterCID)
	if err != nil {
		return err
	}
	if request.Accept {
		err := operations.AcceptRosterRequest(record, request.Reason, requester)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	} else {
		err := operations.RejectRosterRequest(record, request.Reason, requester)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
	}
	return c.JSON(http.StatusOK, model.MakeControllerRosterRequestResponse(record))
}
