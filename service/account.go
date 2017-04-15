package service

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

type tokenResponse struct {
	Token string `json:"token"`
}

func (env *Environment) signIn(context *gin.Context) {
	email := context.Query("email")
	password := context.Query("password")

	account, err := env.Account.FindByEmail(email)
	if err != nil || account.Password != password {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	var token string
	if token, err = env.Token.Create(account.Id); err != nil {
		context.AbortWithStatus(http.StatusBadGateway)
		return
	}
	context.JSON(http.StatusOK, tokenResponse{token})
}
