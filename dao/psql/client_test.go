package psql

import (
	"testing"
)

func TestClientDaoDaoFindById(t *testing.T) {
	dao := ClientDao(db)
	_, err := dao.FindById(1)
	if err != nil {
		t.Error(err)
	}
}
