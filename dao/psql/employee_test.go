package psql

import (
	"testing"
)

func TestEmployeeDaoFindById(t *testing.T) {
	dao := EmployeeDao(db)
	_, err := dao.FindById(1)
	if err != nil {
		t.Error(err)
	}
}

func TestEmployeeDaoFindByEmailAndPassword(t *testing.T) {
	dao := EmployeeDao(db)
	_, err := dao.FindByEmailAndPassword("annin@truck.ru", "18e075e02ee54092bc75d00da7a62f6f")
	if err != nil {
		t.Error(err)
	}
}
