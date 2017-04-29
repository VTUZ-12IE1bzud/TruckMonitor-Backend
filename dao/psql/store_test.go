package psql

import (
	"testing"
)

func TestStoreDaoFindById(t *testing.T) {
	dao := StoreDao(db)
	_, err := dao.FindById(1)
	if err != nil {
		t.Error(err)
	}
}
