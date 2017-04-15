package service

import (
	"gopkg.in/gin-gonic/gin.v1"
	"TruckMonitor-Backend/psql"
	"TruckMonitor-Backend/token"
	"net/http"
)

type AppService struct {
	Port string
	*Environment
}

type DB struct {
	Account psql.AccountDB
}

type Environment struct {
	Token token.TokenParser
	*DB
}

func (service *AppService) Run() {
	// Route
	api := gin.Default()
	v1 := api.Group("/api/v1")
	{
		v1.GET("/login", service.signIn)
	}

	api.NoRoute(service.notFoundError)
	api.Run(service.Port)
}

func (env *Environment) notFoundError(context *gin.Context) {
	context.AbortWithStatus(http.StatusNotFound)
}
