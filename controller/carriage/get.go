package carriage

import (
	"TruckMonitor-Backend/controller/authentication"
	"TruckMonitor-Backend/model"
	"errors"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"net/http"
)

type (
	currentResponse struct {
		Id         int          `json:"id"`
		Carriages  []carriage   `json:"carriages"`
		CheckPoint []checkPoint `json:"checkPoints"`
	}
)
type (
	archiveResponse struct {
		Id          int          `json:"id"`
		CheckPoints []checkPoint `json:"checkPoints"`
	}
)

type (
	onwardResponse struct {
		Id          int          `json:"id"`
		CheckPoints []checkPoint `json:"checkPoints"`
	}
)

func (c *controller) GetCurrent(context *gin.Context) {
	employeeId := context.MustGet(authentication.PARAM_EMPLOYEE_ID).(int)
	dataCarriages, err := c.dao.CarriageDao().FindByDriveAndStatus(employeeId, model.CURRENT)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if len(dataCarriages) == 0 {
		// Если активных рейсов нет
		context.AbortWithStatus(http.StatusNoContent)
		return
	} else if len(dataCarriages) > 1 {
		// Если активных рейсов более 1
		context.AbortWithStatus(http.StatusInternalServerError)
		log.Println(errors.New("len(carriages) > 1"))
		return
	}

	// Если активный рейс 1
	currentCarriage := dataCarriages[0]
	carriageDetails, err := c.dao.CarriageDao().FindDetailById(currentCarriage.Id)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	result := currentResponse{
		Id: currentCarriage.Id,
	}

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
		result.Carriages = append(result.Carriages, carriage{
			Contract:      contract,
			Nomenclatures: nomenclatures,
		})
	}

	result.CheckPoint, err = c.getCheckPoints(currentCarriage.Id)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	context.JSON(http.StatusOK, result)
}

func (c *controller) GetArchive(context *gin.Context) {
	dataCarriages, err := c.dao.CarriageDao().FindByDriveAndStatus(
		context.MustGet(authentication.PARAM_EMPLOYEE_ID).(int), model.ARCHIVE)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	results := make([]archiveResponse, 0)
	for _, data := range dataCarriages {
		checkPoints, err := c.getCheckPoints(data.Id)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		item := archiveResponse{
			Id:          data.Id,
			CheckPoints: checkPoints,
		}
		results = append(results, item)
	}
	context.JSON(http.StatusOK, results)
}

func (c *controller) GetOnward(context *gin.Context) {
	dataCarriages, err := c.dao.CarriageDao().FindByDriveAndStatus(
		context.MustGet(authentication.PARAM_EMPLOYEE_ID).(int), model.ONWARD)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	results := make([]archiveResponse, 0)
	for _, data := range dataCarriages {
		checkPoints, err := c.getCheckPoints(data.Id)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		item := archiveResponse{
			Id:          data.Id,
			CheckPoints: checkPoints,
		}
		results = append(results, item)
	}
	context.JSON(http.StatusOK, results)
}
