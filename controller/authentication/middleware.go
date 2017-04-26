package authentication

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

const HEADER_TOKEN = "X-Auth-Token"
const PARAM_EMPLOYEE_ID = "employeeId"

func (c *controller) Authenticated() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Validate Token
		token := context.Request.Header.Get(HEADER_TOKEN)
		if token == "" {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		employee, err := c.authenticationService.ResolveToken(token)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		context.Set(PARAM_EMPLOYEE_ID, employee.Id)
		context.Next()
	}
}
