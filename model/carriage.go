package model

import "time"

type (
	// Грузоперевозка
	Carriage struct {
		Id        int
		Status    string
		VehicleId int
		DriverId  int
	}

	// Состав груза
	CarriageDetail struct {
		Id         int
		CarriageId int
		ContractId int
	}

	// График движения
	CarriageRoute struct {
		Id           int
		CarriageId   int
		CheckPointId int
		Planned      time.Time
		Fact         time.Time
	}
)
