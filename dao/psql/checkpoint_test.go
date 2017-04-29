package psql

import (
	"testing"
)

func TestCheckPointDaoFindById(t *testing.T) {
	dao := CheckPointDao(db)
	_, err := dao.FindById(1)
	if err != nil {
		t.Error(err)
	}
}
