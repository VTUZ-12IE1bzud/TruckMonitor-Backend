package psql

import (
	"TruckMonitor-Backend/dao"
	"TruckMonitor-Backend/model"
	"database/sql"
	"log"
)

type psqlCarriageDao struct {
	context PsqlContext
}

func CarriageDao(context PsqlContext) dao.CarriageDao {
	return &psqlCarriageDao{context}
}

func (dao *psqlCarriageDao) db() *sql.DB {
	return dao.context.GetDb()
}

func (dao *psqlCarriageDao) FindByDriveAndStatus(driverId int, status string) ([]*model.Carriage, error) {
	rows, err := dao.db().Query("SELECT * FROM carriage WHERE driver_id=$1 AND status=$2", driverId, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := make([]*model.Carriage, 0)
	for rows.Next() {
		item := new(model.Carriage)
		err := rows.Scan(&item.Id, &item.Status, &item.VehicleId, &item.DriverId)
		if err != nil {
			log.Println(err)
		} else {
			items = append(items, item)
		}
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	return items, nil
}

func (dao *psqlCarriageDao) FindDetailById(carriageId int) ([]*model.CarriageDetail, error) {
	rows, err := dao.db().Query("SELECT * FROM carriage_detail WHERE carriage_id=$1", carriageId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	items := make([]*model.CarriageDetail, 0)
	for rows.Next() {
		item := new(model.CarriageDetail)
		err := rows.Scan(&item.Id, &item.CarriageId, &item.ContractId)
		if err != nil {
			log.Println(err)
		} else {
			items = append(items, item)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (dao *psqlCarriageDao) FindRouteByCarriage(carriageId int) ([]*model.CarriageRoute, error) {
	rows, err := dao.db().Query("SELECT * FROM carriage_route WHERE carriage_id=$1", carriageId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	items := make([]*model.CarriageRoute, 0)
	for rows.Next() {
		item := new(model.CarriageRoute)
		err := rows.Scan(&item.Id, &item.CarriageId, &item.CheckPointId, &item.Planned, &item.Fact)
		if err != nil {
			log.Println(err)
		} else {
			items = append(items, item)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
