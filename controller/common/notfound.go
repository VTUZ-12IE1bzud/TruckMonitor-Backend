package common

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func (ctx *commonController) NotFound(context *gin.Context) {
	context.AbortWithStatus(http.StatusNotFound)
}
