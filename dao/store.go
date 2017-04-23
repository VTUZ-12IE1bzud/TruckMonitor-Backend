package dao

import "TruckMonitor-Backend/model"

type StoreDao interface {
	FindById(id int) (*model.Store, error)
}
