package dao

import "TruckMonitor-Backend/model"

type CheckPointDao interface {
	FindById(id int) (*model.CheckPoint, error)
}