package controller

import (
	"TruckMonitor-Backend/context"
	"TruckMonitor-Backend/controller/authentication"
	"TruckMonitor-Backend/controller/common"
	"gopkg.in/gin-gonic/gin.v1"
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

	// Route
	api := gin.Default()
	v1 := api.Group("/api/v1")
	{
		v1.GET("/login", authenticationController.Get)
		v1.Use(authenticationController.Authenticated())
		{

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
