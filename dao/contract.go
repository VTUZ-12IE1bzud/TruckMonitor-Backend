package dao

import "TruckMonitor-Backend/model"

type ContractDao interface {
	FindById(id int) (*model.Contract, error)
	FindDetails(contractId int) ([]*model.ContractDetail, error)
}
