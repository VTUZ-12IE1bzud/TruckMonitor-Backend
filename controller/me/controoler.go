package me

import "TruckMonitor-Backend/dao"

type controller struct {
	employeeDao dao.EmployeeDao
}

func Controller(employeeDao dao.EmployeeDao) *controller {
	return &controller{
		employeeDao: employeeDao,
	}
}
