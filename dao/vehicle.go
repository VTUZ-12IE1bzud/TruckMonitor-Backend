package dao

import "TruckMonitor-Backend/model"

type VehicleDao interface {
	FindTypeById(id int) (*model.VehicleType, error)
	FindBrandById(id int) (*model.VehicleBrand, error)
	FindModelById(id int) (*model.VehicleModel, error)
	FindById(id int) (*model.Vehicle, error)
}
