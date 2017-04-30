package controller

import (
	"TruckMonitor-Backend/context"
	"TruckMonitor-Backend/controller/authentication"
	"TruckMonitor-Backend/controller/carriage"
	"TruckMonitor-Backend/controller/common"
	"gopkg.in/gin-gonic/gin.v1"
	"TruckMonitor-Backend/controller/checkpoint"
)

type Router interface {
	Run(port string) error
}

type router struct {
	context context.ApplicationContext
}

func (r router) Run(port string) error {
	// Controllers
	commonController := common.Controller()
	authenticationController := authentication.Controller(r.context.ServiceContext().AuthenticationService())
	carriageController := carriage.Controller(r.context.DaoContext())
	checkPointController := checkpoint.Controller(r.context.DaoContext().CarriageDao())

	// Route
	api := gin.Default()
	v1 := api.Group("/api/v1")
	{
		v1.GET("/login", authenticationController.GetToken)
		v1.Use(authenticationController.Authenticated())
		{
			v1.GET("/carriage/current", carriageController.GetCurrent)
			v1.GET("/carriage/archive", carriageController.GetArchive)
			v1.GET("/carriage/onward", carriageController.GetOnward)
			v1.POST("/checkpoint/:checkpoint", checkPointController.CreateFactTimestamp)
		}
	}
	api.NoRoute(commonController.NotFound)
	return api.Run(port)
}

func NewRouter(context context.ApplicationContext) Router {
	return &router{
		context: context,
	}
}
