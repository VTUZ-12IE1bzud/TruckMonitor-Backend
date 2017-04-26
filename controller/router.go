package controller

import (
	"TruckMonitor-Backend/context"
	"gopkg.in/gin-gonic/gin.v1"
	"TruckMonitor-Backend/controller/common"
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

	// Route
	api := gin.Default()
	//v1 := api.Group("/api/v1")
	{
		//v1.GET("/login", service.signIn)
		//v1.Use(service.authRequiredMiddleware())
		//{
		//	freightTraffic := v1.Group("/freighttraffic")
		//	freightTraffic.GET("/current", service.currentCargoTransportation)
		//}
	}
	api.NoRoute(commonController.NotFound)
	return api.Run(port)
}

func NewRouter(context context.ApplicationContext) Router {
	return &router{
		context: context,
	}
}
