package checkpoint

import (
	"TruckMonitor-Backend/controller/authentication"
	"TruckMonitor-Backend/model"
	"errors"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (c *controller) CreateFactTimestamp(context *gin.Context) {
	checkPointId, err := strconv.Atoi(context.Param("checkpoint"))
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
		return
	}
	employeeId := context.MustGet(authentication.PARAM_EMPLOYEE_ID).(int)
	dataCarriages, err := c.carriageDao.FindByDriveAndStatus(employeeId, model.CURRENT)
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
		return
	}

	if len(dataCarriages) == 0 {
		// Если активных рейсов нет
		context.AbortWithStatus(http.StatusBadRequest)
		return
	} else if len(dataCarriages) > 1 {
		// Если активных рейсов более 1
		context.AbortWithStatus(http.StatusInternalServerError)
		log.Println(errors.New("len(carriages) > 1"))
		return
	}

	// Если активный рейс 1
	currentCarriage := dataCarriages[0]
	c.carriageDao.CreateFactTimestamp(currentCarriage.Id, checkPointId, time.Now())
	context.AbortWithStatus(http.StatusOK)
}
