package api

import (
	"fmt"
	"github.com/VATUSA/api-v3/pkg/auth"
	"github.com/VATUSA/api-v3/pkg/controller"
	db "github.com/VATUSA/api-v3/pkg/database"
	"github.com/VATUSA/api-v3/pkg/datamodel/response"
	"github.com/VATUSA/api-v3/pkg/facility"
	"github.com/VATUSA/api-v3/pkg/roster"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetFacilityRoster godoc
//	@Summary	Get Facility Roster
//	@Tags		facility
//	@Param		facility	path	int	true	"Facility"
//	@Router		/v3/facility/{id}/roster [get]
//	@Success	200	{array}		response.ControllerLimited
//	@Success	200	{array}		response.Controller
//	@Failure	400	{object}	response.ErrorMessage
func GetFacilityRoster(c echo.Context) error {
	fac := facility.Facility(c.Param("facility"))
	controllers, err := db.FetchControllersByHomeFacility(fac)
	visitors, err := db.FetchControllersByVisitingFacility(fac)
	controllers = append(controllers, visitors...)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to fetch data")
	}
	if auth.CanReadControllerSensitiveData(c) {
		var output []*response.Controller
		for _, v := range controllers {
			output = append(output, response.MakeController(&v))
		}
		return c.JSON(http.StatusOK, output)
	} else if auth.IsAuthenticated(c) {
		var output []*response.Controller
		for _, v := range controllers {
			output = append(output, response.MakeControllerRedacted(&v))
		}
		return c.JSON(http.StatusOK, output)
	} else {
		var output []*response.ControllerLimited
		for _, v := range controllers {
			output = append(output, response.MakeControllerLimited(&v))
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
	fac := facility.Facility(c.Param("facility"))
	var request RemoveFromRosterRequest
	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Payload is incorrectly formatted")
	}
	controllerModel, err := db.FetchControllerByCID(request.CID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("CID %d does not exist", request.CID))
	}
	var requester *db.Controller
	if request.RequesterCID != nil {
		requester, err = db.FetchControllerByCID(*request.RequesterCID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("CID %d does not exist", request.RequesterCID))
		}
	}
	if controllerModel.Facility == fac {
		// Home Removal
		err := controller.RemoveFromFacility(controllerModel, requester, request.Reason)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	} else if controller.IsVisiting(controllerModel, fac) {
		err := controller.RemoveVisitor(controllerModel, fac, requester, request.Reason)
		if err != nil {

		}
	} else {
		return echo.NewHTTPError(
			http.StatusBadRequest, fmt.Sprintf("Controller %d is not on the %s roster", request.CID, fac))
	}
	return nil
}

func GetPendingRosterRequests(c echo.Context) error {
	fac := c.Param("facility")
	requests, err := db.FetchPendingRequestsByFacility(fac)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	var output []*response.ControllerRosterRequest
	for _, v := range requests {
		output = append(output, response.MakeControllerRosterRequestResponse(&v))
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
	fac := facility.Facility(c.Param("facility"))
	var request ProcessRosterRequestRequest
	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Payload is incorrectly formatted")
	}
	record, err := db.FetchRequestById(request.ID)
	if record.Status != roster.StatusPending {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Roster Request %d is not Pending", request.ID))
	}
	if record.Facility != fac {
		return echo.NewHTTPError(
			http.StatusBadRequest, fmt.Sprintf("Roster Request %d is not for facility %s", request.ID, fac))
	}
	requester, err := db.FetchControllerByCID(request.RequesterCID)
	if err != nil {
		return err
	}
	if request.Accept {
		err := controller.AcceptRosterRequest(record, request.Reason, requester)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	} else {
		err := controller.RejectRosterRequest(record, request.Reason, requester)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
	}
	return c.JSON(http.StatusOK, response.MakeControllerRosterRequestResponse(record))
}
