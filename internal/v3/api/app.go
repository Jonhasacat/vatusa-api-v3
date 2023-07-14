package api

import (
	middleware3 "github.com/VATUSA/api-v3/internal/v3/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func App() *echo.Echo {
	e := echo.New()
	e.Use(middleware3.AuthContext) // AuthContext needs to be first as it adds data to Context
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} - HTTP ${status}\n",
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	gController := e.Group("/v3/controller/:cid")
	gController.GET("", GetController)
	gController.GET("/training/record", GetControllerTrainingRecords)

	gFacility := e.Group("/v3/facility/:facility")

	// Roster
	gFacility.GET("/roster", GetFacilityRoster)
	gFacility.POST("/roster/remove", RemoveFromRoster, middleware3.AuthFacilityATMOrDATMOrToken)
	gFacility.GET("/roster/request", GetPendingRosterRequests, middleware3.AuthFacilityATMOrDATMOrToken)
	gFacility.POST("/roster/request", ProcessRosterRequest, middleware3.AuthFacilityATMOrDATMOrToken)

	// Solo
	gFacility.GET("/solo", GetSoloCertificationsByFacility)
	gFacility.POST("/solo", CreateSoloCertification, middleware3.AuthFacilityInstructorOrToken)
	gFacility.PUT("/solo/:id", ModifySoloCertification, middleware3.AuthFacilityInstructorOrToken)
	gFacility.DELETE("/solo/:id", DeleteSoloCertification, middleware3.AuthFacilityInstructorOrToken)

	// Training Records
	gFacility.GET("/training/record", GetFacilityTrainingRecords,
		middleware3.AuthFacilityTrainingStaffOrToken)
	gFacility.POST("/training/record", CreateTrainingRecord, middleware3.AuthFacilityTrainingStaffOrToken)
	gFacility.PUT("/training/record/:id", ModifyTrainingRecord, middleware3.AuthFacilityTrainingStaffOrToken)
	gFacility.DELETE("/training/record/:id", DeleteTrainingRecord, middleware3.AuthFacilitySeniorStaffOrToken)

	//gPublic := e.Group("/v3/public")

	gSelf := e.Group("/v3/self")
	gSelf.GET("", GetMyInfo, middleware3.AuthControllerOnly)

	return e
}
