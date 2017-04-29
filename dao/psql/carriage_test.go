package psql

import (
	"TruckMonitor-Backend/model"
	"testing"
)

func TestCarriageDaoFindByDriveAndStatus(t *testing.T) {
	dao := CarriageDao(db)
	_, err := dao.FindByDriveAndStatus(4, model.CURRENT)
	if err != nil {
		t.Error(err)
	}
}

func TestCarriageDaoFindDetailById(t *testing.T) {
	dao := CarriageDao(db)
	_, err := dao.FindDetailById(1)
	if err != nil {
		t.Error(err)
	}
}

func TestCarriageDaoFindRouteByCarriage(t *testing.T) {
	dao := CarriageDao(db)
	_, err := dao.FindRouteByCarriage(1)
	if err != nil {
		t.Error(err)
	}
}
