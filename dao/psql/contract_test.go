package psql

import (
	"testing"
)

func TestContractDaoFindById(t *testing.T) {
	dao := ContractDao(db)
	_, err := dao.FindById(1)
	if err != nil {
		t.Error(err)
	}
}

func TestContractDaoFindDetails(t *testing.T) {
	dao := ContractDao(db)
	_, err := dao.FindDetails(1)
	if err != nil {
		t.Error(err)
	}
}
