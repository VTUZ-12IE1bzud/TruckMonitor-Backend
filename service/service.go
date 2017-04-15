package service

import (
	"gopkg.in/gin-gonic/gin.v1"
	"TruckMonitor-Backend/psql"
	"net/http"
)

type AppService struct {
	Port string
	*Environment

}

type Environment struct {
	Account psql.AccountDB
}

func (service *AppService) Run() {
	// Route
	api := gin.Default()

	api.NoRoute(service.notFoundError)
	api.Run(service.Port)
}

func (env *Environment) notFoundError(context *gin.Context) {
	context.AbortWithStatus(http.StatusNotFound)
}