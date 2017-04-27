package carriage

import (
	"TruckMonitor-Backend/model"
	"time"
	"fmt"
	"TruckMonitor-Backend/controller/authentication"
	"log"
	"errors"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

type (
	currentResponse struct {
		Id         int `json:"id"`
		Carriages  []Carriage `json:"carriages"`
		CheckPoint []*checkPoint `json:"checkPoints"`
	}

	Carriage struct {
		Contract      *Contract `json:"contract"`
		Nomenclatures []*Nomenclature `json:"nomenclatures"`
	}

	Contract struct {
		Number      string `json:"number"`
		CustomsLink string `json:"customsLink"`
		From        Store `json:"from"`
		Before      Store `json:"before"`
		Packaging   string `json:"packaging"`
		Manager     *manager `json:"manager"`
	}

	Store struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	}

	Nomenclature struct {
		Name    string  `json:"name"`
		Amount  float32 `json:"amount"`
		Measure string  `json:"measure"`
	}

	manager struct {
		Name      string `json:"name"`
		PhotoLink string `json:"photolink"`
		Phone     string `json:"phone"`
	}

	checkPoint struct {
		Coordinates *coordinates `json:"coordinates"`
		Planned     string `json:"planned"`
		Fact        string `json:"fact"`
	}

	coordinates struct {
		Latitude  float32 `json:"latitude"`
		Longitude float32 `json:"longitude"`
	}
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
	carriage := carriages[0]
	carriageDetails, err := c.dao.CarriageDao().FindDetailById(carriage.Id)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	items := make([]Carriage, 0)
	for _, carriageDetail := range carriageDetails {
		contract, err := c.getContract(carriageDetail.ContractId)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		nomenclatures, err := c.createNomenclature(carriageDetail.ContractId)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		item := Carriage{
			Contract:      contract,
			Nomenclatures: nomenclatures,
		}
		items = append(items, item)
	}

	checkPoints, err := c.getCheckPoints(carriage.Id)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	context.JSON(http.StatusOK, currentResponse{
		Id:         carriage.Id,
		Carriages:  items,
		CheckPoint: checkPoints,
	})
}

func (c *controller) getContract(contractId int) (*Contract, error) {
	contract, err := c.dao.ContractDao().FindById(contractId)
	if err != nil {
		return nil, err
	}

	from, err := c.getStore(contract.StoreFromId)
	if err != nil {
		return nil, err
	}
	before, err := c.getStore(contract.StoreBeforeId)
	if err != nil {
		return nil, err
	}

	manager, err := c.getManager(contract.ManagerId)
	if err != nil {
		return nil, err
	}

	packaging, err := c.getPackaging(contract.PackagingId)
	if err != nil {
		return nil, err
	}

	return &Contract{
		Number:      contract.Number,
		CustomsLink: contract.ConfirmationCustomsLink,
		From:        *from,
		Before:      *before,
		Manager:     manager,
		Packaging:   packaging,
	}, nil
}

func (c *controller) getStore(storeId int) (*Store, error) {
	store, err := c.dao.StoreDao().FindById(storeId)
	if err != nil {
		return nil, err
	}
	return &Store{
		Name:    store.Name,
		Address: store.Address,
	}, nil
}

func (c *controller) getPackaging(packagingId int) (string, error) {
	packaging, err := c.dao.PackagingDao().FindById(packagingId)
	if err != nil {
		return "", err
	}
	return packaging.Name, nil
}

func (c *controller) createNomenclature(contractId int) ([]*Nomenclature, error) {
	contracts, err := c.dao.ContractDao().FindDetails(contractId)
	if err != nil {
		return nil, err
	}
	items := make([]*Nomenclature, 0)
	for _, contract := range contracts {
		nomenclature, err := c.getNomenclature(contract.NomenclatureId)
		if err != nil {
			return nil, err
		}
		measure, err := c.getMeasure(nomenclature.MeasureId)

		item := &Nomenclature{
			Name:    nomenclature.Name,
			Measure: measure,
			Amount:  contract.Amount,
		}
		items = append(items, item)
	}
	return items, nil
}

func (c *controller) getNomenclature(nomenclatureId int) (*model.Nomenclature, error) {
	nomenclature, err := c.dao.NomenclatureDao().FindById(nomenclatureId)
	if err != nil {
		return nil, err
	}
	return nomenclature, nil
}

func (c *controller) getMeasure(measureId int) (string, error) {
	measure, err := c.dao.MeasureDao().FindById(measureId)
	if err != nil {
		return "", err
	}
	return measure.Abbreviation, nil
}

func (c *controller) getManager(managerId int) (*manager, error) {
	employee, err := c.dao.EmployeeDao().FindById(managerId)
	if err != nil {
		return nil, err
	}
	return &manager{
		Name:      convertEmployeeName(employee),
		PhotoLink: employee.Photo,
		Phone:     employee.Phone,
	}, nil
}

func (c *controller) getCheckPoints(carriageId int) ([]*checkPoint, error) {
	routes, err := c.dao.CarriageDao().FindRouteByCarriage(carriageId)
	if err != nil {
		return nil, err
	}
	items := make([]*checkPoint, 0)
	for _, route := range routes {
		coordinates, err := c.getCoordinates(route.CheckPointId)
		if err != nil {
			return nil, err
		}
		item := &checkPoint{
			Coordinates: coordinates,
			Planned:     convertDateTime(route.Planned),
			Fact:        convertDateTime(route.Fact),
		}
		items = append(items, item)

	}
	return items, nil
}

func (c *controller) getCoordinates(checkPointId int) (*coordinates, error) {
	checkPoint, err := c.dao.CheckPointDao().FindById(checkPointId)
	if err != nil {
		return nil, err
	}
	return &coordinates{
		Longitude: checkPoint.Longitude,
		Latitude:  checkPoint.Latitude,
	}, nil
}

func convertDateTime(dateTime time.Time) string {
	return dateTime.Format("2014-11-17 23:02:03 +0000 UTC")
}

func convertEmployeeName(employee *model.Employee) string {
	return fmt.Sprintf("%s %s %s", employee.Surname, employee.Name, employee.Patronymic)
}
