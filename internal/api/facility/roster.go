package facility

import (
	"fmt"
	"github.com/VATUSA/api-v3/internal/core"
	"github.com/VATUSA/api-v3/internal/database"
	"github.com/VATUSA/api-v3/pkg/constants"
	"github.com/VATUSA/api-v3/pkg/facility_api/model"
	"github.com/VATUSA/api-v3/pkg/facility_api/translator"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetRoster(c echo.Context) error {
	facilityID := c.Get("facility").(string)
	controllers, err := database.FetchControllersByHomeFacility(facilityID)
	if err != nil {
		return err
	}
	visitors, err := database.FetchControllersByVisitingFacility(facilityID)
	if err != nil {
		return err
	}
	controllers = append(controllers, visitors...)
	output := translator.TranslateControllers(controllers)
	return c.JSON(http.StatusOK, output)
}

func RemoveFromRoster(c echo.Context) error {
	facilityID := c.Get("facility").(string)
	var request model.RosterRemoveRequest
	err := c.Bind(request)
	if err != nil {
		return ErrorBadPayload
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
	if controller.Facility == facilityID {
		// Home Removal
		err := core.RemoveFromFacility(controller, requester, request.Reason)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	} else if core.IsVisiting(controller, facilityID) {
		err := core.RemoveVisitor(controller, facilityID, requester, request.Reason)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	} else {
		return echo.NewHTTPError(
			http.StatusBadRequest, fmt.Sprintf("Controller %d is not on the %s roster", request.CID, facilityID))
	}
	return c.JSON(http.StatusOK, nil)
}

func GetPendingRosterRequests(c echo.Context) error {
	facilityID := c.Get("facility").(string)
	requests, err := database.FetchPendingRequestsByFacility(facilityID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	output := translator.TranslateRosterRequests(requests)
	return c.JSON(http.StatusOK, output)
}

func ProcessRosterRequest(c echo.Context) error {
	facilityID := c.Get("facility").(string)
	var request model.ProcessRosterRequestRequest
	err := c.Bind(request)
	if err != nil {
		return ErrorBadPayload
	}
	record, err := database.FetchRequestById(request.ID)
	if record.Status != constants.StatusPending {
		return ErrorBadRecordState
	}
	if record.Facility != facilityID {
		return ErrorRequestFacilityMismatch
	}
	requester, err := database.FetchControllerByCID(request.RequesterCID)
	if err != nil {
		return err
	}
	if request.Accept {
		err = core.AcceptRosterRequest(record, request.Reason, requester)
	} else {
		err = core.RejectRosterRequest(record, request.Reason, requester)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, nil)
}
