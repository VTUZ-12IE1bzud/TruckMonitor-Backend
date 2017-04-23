package dao

import "TruckMonitor-Backend/model"

type ClientDao interface {
	FindById(id int) (*model.Client, error)
}
