package authentication

import (
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"net/http"
)

type (
	response struct {
		Token string `json:"token"`
	}
)

func (c *controller) GetToken(context *gin.Context) {
	email := context.Query("email")
	password := context.Query("password")

	token, err := c.authenticationService.CreateToken(email, password)
	if err != nil {
		log.Println(err)
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}
	context.JSON(http.StatusOK, response{
		Token: token,
	})
}
