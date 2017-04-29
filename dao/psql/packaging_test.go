package psql

import (
	"testing"
)

func TestPackagingDaoFindById(t *testing.T) {
	dao := PackagingDao(db)
	_, err := dao.FindById(1)
	if err != nil {
		t.Error(err)
	}
}
