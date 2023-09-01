package facility

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	ErrorBadPayload              = echo.NewHTTPError(http.StatusBadRequest, "payload is incorrectly formatted")
	ErrorBadRecordState          = echo.NewHTTPError(http.StatusBadRequest, "operation is not permitted due to record status")
	ErrorRequestFacilityMismatch = echo.NewHTTPError(http.StatusForbidden, "facility mismatch")
)
