package psql

import (
	"testing"
)

func TestNomenclatureDaoFindById(t *testing.T) {
	dao := NomenclatureDao(db)
	_, err := dao.FindById(1)
	if err != nil {
		t.Error(err)
	}
}
