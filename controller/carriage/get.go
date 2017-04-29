package carriage

import (
	"TruckMonitor-Backend/controller/authentication"
	"TruckMonitor-Backend/model"
	"errors"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"net/http"
)

func (c *controller) GetCurrent(context *gin.Context) {
	employeeId := context.MustGet(authentication.PARAM_EMPLOYEE_ID).(int)
	carriages, err := c.dao.CarriageDao().FindByDriveAndStatus(employeeId, model.CURRENT)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if len(carriages) == 0 {
		// Если активных рейсов нет
		context.AbortWithStatus(http.StatusNoContent)
		return
	} else if len(carriages) > 1 {
		// Если активных рейсов более 1
		context.AbortWithStatus(http.StatusInternalServerError)
		log.Println(errors.New("len(carriages) > 1"))
		return
	}

	// Если активный рейс 1
	currentCarriage := carriages[0]
	carriageDetails, err := c.dao.CarriageDao().FindDetailById(currentCarriage.Id)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	items := make([]carriage, 0)
	for _, carriageDetail := range carriageDetails {
		contract, err := c.getContract(carriageDetail.ContractId)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		nomenclatures, err := c.getNomenclature(carriageDetail.ContractId)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		item := carriage{
			Contract:      contract,
			Nomenclatures: nomenclatures,
		}
		items = append(items, item)
	}

	checkPoints, err := c.getCheckPoints(currentCarriage.Id)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	context.JSON(http.StatusOK, currentResponse{
		Id:         currentCarriage.Id,
		Carriages:  items,
		CheckPoint: checkPoints,
	})
}








