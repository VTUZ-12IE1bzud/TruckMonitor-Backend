package psql

import (
	"testing"
)

func TestVehicleDaoFindTypeById(t *testing.T) {
	dao := VehicleDao(db)
	_, err := dao.FindById(1)
	if err != nil {
		t.Error(err)
	}
}

func TestVehicleDaoFindBrandById(t *testing.T) {
	dao := VehicleDao(db)
	_, err := dao.FindBrandById(1)
	if err != nil {
		t.Error(err)
	}
}

func TestVehicleDaoFindModelById(t *testing.T) {
	dao := VehicleDao(db)
	_, err := dao.FindModelById(1)
	if err != nil {
		t.Error(err)
	}
}

func TestVehicleDaoFindById(t *testing.T) {
	dao := VehicleDao(db)
	_, err := dao.FindById(1)
	if err != nil {
		t.Error(err)
	}
}
