package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	middleware2 "vatusa-api-v3/v3/api/middleware"
	"vatusa-api-v3/v3/api/routes"
)

func App() *echo.Echo {
	e := echo.New()
	e.Use(middleware2.AuthContext) // AuthContext needs to be first as it adds data to Context
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} - HTTP ${status}\n",
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	gController := e.Group("/v3/controller/:cid")
	gController.GET("", routes.GetController)
	gController.GET("/training/record", routes.GetControllerTrainingRecords)

	gFacility := e.Group("/v3/facility/:facility")

	// Roster
	gFacility.GET("/roster", routes.GetFacilityRoster)
	gFacility.POST("/roster/remove", routes.RemoveFromRoster, middleware2.AuthFacilityATMOrDATMOrToken)
	gFacility.GET("/roster/request", routes.GetPendingRosterRequests, middleware2.AuthFacilityATMOrDATMOrToken)
	gFacility.POST("/roster/request", routes.ProcessRosterRequest, middleware2.AuthFacilityATMOrDATMOrToken)

	// Solo
	gFacility.GET("/solo", routes.GetSoloCertificationsByFacility)
	gFacility.POST("/solo", routes.CreateSoloCertification, middleware2.AuthFacilityInstructorOrToken)
	gFacility.PUT("/solo/:id", routes.ModifySoloCertification, middleware2.AuthFacilityInstructorOrToken)
	gFacility.DELETE("/solo/:id", routes.DeleteSoloCertification, middleware2.AuthFacilityInstructorOrToken)

	// Training Records
	gFacility.GET("/training/record", routes.GetFacilityTrainingRecords,
		middleware2.AuthFacilityTrainingStaffOrToken)
	gFacility.POST("/training/record", routes.CreateTrainingRecord, middleware2.AuthFacilityTrainingStaffOrToken)
	gFacility.PUT("/training/record/:id", routes.ModifyTrainingRecord, middleware2.AuthFacilityTrainingStaffOrToken)
	gFacility.DELETE("/training/record/:id", routes.DeleteTrainingRecord, middleware2.AuthFacilitySeniorStaffOrToken)

	//gPublic := e.Group("/v3/public")

	gSelf := e.Group("/v3/self")
	gSelf.GET("", routes.GetMyInfo, middleware2.AuthControllerOnly)

	return e
}
