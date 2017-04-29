package psql

import (
	"testing"
)

func TestMeasureDaoFindById(t *testing.T) {
	dao := MeasureDao(db)
	_, err := dao.FindById(1)
	if err != nil {
		t.Error(err)
	}
}
