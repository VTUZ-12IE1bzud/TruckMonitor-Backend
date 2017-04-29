package carriage

import (
	"TruckMonitor-Backend/context"
	"TruckMonitor-Backend/model"
)

type controller struct {
	dao context.DaoContext
}

func Controller(dao context.DaoContext) *controller {
	return &controller{
		dao: dao,
	}
}

func (c *controller) getContract(contractId int) (result contract, err error) {
	contract, err := c.dao.ContractDao().FindById(contractId)
	if err != nil {
		return
	}

	result.Number = contract.Number
	result.CustomsLink = contract.ConfirmationCustomsLink

	if result.From, err = c.getStore(contract.StoreFromId); err != nil {
		return
	}
	if 	result.Before, err = c.getStore(contract.StoreBeforeId); err != nil {
		return
	}
	if 	result.Manager, err = c.getManager(contract.ManagerId); err != nil {
		return
	}
	if 	result.Packaging, err = c.getPackaging(contract.PackagingId); err != nil {
		return
	}
	return
}

func (c *controller) getStore(storeId int) (result store, err error) {
	store, err := c.dao.StoreDao().FindById(storeId)
	if err != nil {
		return
	}
	result.Name = store.Name
	result.Address = store.Address
	return
}



func (c *controller) getPackaging(packagingId int) (result string, err error) {
	packaging, err := c.dao.PackagingDao().FindById(packagingId)
	if err != nil {
		return
	}
	result = packaging.Name
	return
}

func (c *controller) getNomenclature(contractId int) (result []nomenclature, err error) {
	contracts, err := c.dao.ContractDao().FindDetails(contractId)
	if err != nil {
		return
	}
	for _, contract := range contracts {
		var nom *model.Nomenclature
		nom, err = c.dao.NomenclatureDao().FindById(contract.NomenclatureId)
		if err != nil {
			return
		}
		var measure string
		measure, err = c.getMeasure(nom.MeasureId)

		item := nomenclature{
			Name:    nom.Name,
			Measure: measure,
			Amount:  contract.Amount,
		}
		result = append(result, item)
	}
	return
}

func (c *controller) getMeasure(measureId int) (result string, err error) {
	measure, err := c.dao.MeasureDao().FindById(measureId)
	if err != nil {
		return
	}
	result = measure.Abbreviation
	return
}

func (c *controller) getManager(managerId int) (result manager, err error) {
	employee, err := c.dao.EmployeeDao().FindById(managerId)
	if err != nil {
		return
	}
	result = manager{
		Name:      convertEmployeeName(employee),
		PhotoLink: employee.Photo,
		Phone:     employee.Phone,
	}
	return
}

func (c *controller) getCheckPoints(carriageId int) (result []checkPoint, err error) {
	routes, err := c.dao.CarriageDao().FindRouteByCarriage(carriageId)
	if err != nil {
		return
	}
	for _, route := range routes {
		var coordinate coordinates
		coordinate, err = c.getCoordinates(route.CheckPointId)
		if err != nil {
			return
		}
		item := checkPoint{
			Coordinates: coordinate,
			Planned:     route.Planned.UTC(),
		}
		if route.Fact.IsZero() {
			item.Fact = nil
		} else {
			var time = route.Fact.UTC(); item.Fact = &time
		}
		result = append(result, item)
	}
	return
}

func (c *controller) getCoordinates(checkPointId int) (result coordinates, err error) {
	checkPoint, err := c.dao.CheckPointDao().FindById(checkPointId)
	if err != nil {
		return
	}
	result = coordinates{
		Longitude: checkPoint.Longitude,
		Latitude:  checkPoint.Latitude,
	}
	return
}
