package dao

import "TruckMonitor-Backend/model"

type MeasureDao interface {
	FindById(id int) (*model.Measure, error)
}
