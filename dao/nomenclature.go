package dao

import "TruckMonitor-Backend/model"

type NomenclatureDao interface {
	FindById(id int) (*model.Nomenclature, error)
}
