package dao

import "TruckMonitor-Backend/model"

type EmployeeDao interface {
	FindById(id int) (*model.Employee, error)
	FindByEmailAndPassword(email string, password string) (*model.Employee, error)
}
