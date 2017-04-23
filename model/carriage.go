package model

import "time"

// Статус грузоперевозки
const (
	CURRENT string = "current"
	ARCHIVE string = "archive"
	ONWARD  string = "onward"
)

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
