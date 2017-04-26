package psql

import (
	"TruckMonitor-Backend/dao"
	"database/sql"
	"TruckMonitor-Backend/model"
)

type psqlVehicle struct {
	context PsqlContext
}

func VehicleDao(context PsqlContext) dao.VehicleDao {
	return &psqlVehicle{context}
}

func (dao *psqlVehicle) db() *sql.DB {
	return dao.context.GetDb()
}

func (dao *psqlVehicle) FindTypeById(id int) (*model.VehicleType, error) {
	var data model.VehicleType
	row := dao.db().QueryRow("SELECT * FROM vehicle_type WHERE id=$1", id)
	err := row.Scan(&data.Id, &data.Name)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (dao *psqlVehicle) FindBrandById(id int) (*model.VehicleBrand, error) {
	var data model.VehicleBrand
	row := dao.db().QueryRow("SELECT * FROM vehicle_brand WHERE id=$1", id)
	err := row.Scan(&data.Id, &data.TypeId, &data.Name)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (dao *psqlVehicle) FindModelById(id int) (*model.VehicleModel, error) {
	var data model.VehicleModel
	row := dao.db().QueryRow("SELECT * FROM vehicle_model WHERE id=$1", id)
	err := row.Scan(&data.Id, &data.BrandId, &data.Name, &data.CapacityMax)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (dao *psqlVehicle) FindById(id int) (*model.Vehicle, error) {
	var data model.Vehicle
	row := dao.db().QueryRow("SELECT * FROM vehicle WHERE id=$1", id)
	err := row.Scan(&data.Id, &data.ModelId, &data.Vin, &data.LicencePlate, &data.RegionNumber)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
