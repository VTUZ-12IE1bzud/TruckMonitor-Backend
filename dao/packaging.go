package dao

import "TruckMonitor-Backend/model"

type PackagingDao interface {
	FindById(id int) (*model.Packaging, error)
}
