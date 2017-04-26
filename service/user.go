package service

import (
	"TruckMonitor-Backend/model"
	"TruckMonitor-Backend/dao"
	"crypto/md5"
	"encoding/hex"
)

type (
	UserService interface {
		Get(id int) (*model.Employee, error)
		Validate(email string, password string) (*model.Employee, error)
	}

	userService struct {
		employeeDao dao.EmployeeDao
	}
)

func NewUserService(employeeDao dao.EmployeeDao) UserService {
	return &userService{
		employeeDao: employeeDao,
	}
}


func (s userService) Get(id int) (*model.Employee, error) {
	return s.employeeDao.FindById(id)
}


func (s userService) Validate(email string, password string) (*model.Employee, error) {
	return s.employeeDao.FindByEmailAndPassword(email, computeHash(password))
}

func computeHash(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}
