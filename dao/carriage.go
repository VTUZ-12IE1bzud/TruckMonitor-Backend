package dao

import (
	"TruckMonitor-Backend/model"
	"time"
)

type CarriageDao interface {
	FindByDriveAndStatus(driverId int, status string) ([]*model.Carriage, error)
	FindDetailById(carriageId int) ([]*model.CarriageDetail, error)
	FindRouteByCarriage(carriageId int) ([]*model.CarriageRoute, error)
	CreateFactTimestamp(carriageId int, checkPointId int, time time.Time) error
}
