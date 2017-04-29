package psql

import (
	"TruckMonitor-Backend/dao"
	"TruckMonitor-Backend/model"
	"database/sql"
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
		if err := rows.Scan(&item.Id, &item.Status, &item.VehicleId, &item.DriverId); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (dao *psqlCarriageDao) FindDetailById(carriageId int) ([]*model.CarriageDetail, error) {
	rows, err := dao.db().Query("SELECT * FROM carriage_detail WHERE carriage_id=$1", carriageId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := make([]*model.CarriageDetail, 0)
	for rows.Next() {
		item := new(model.CarriageDetail)
		if err := rows.Scan(&item.Id, &item.CarriageId, &item.ContractId); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (dao *psqlCarriageDao) FindRouteByCarriage(carriageId int) ([]*model.CarriageRoute, error) {
	rows, err := dao.db().Query("SELECT * FROM carriage_route WHERE carriage_id=$1", carriageId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := make([]*model.CarriageRoute, 0)
	for rows.Next() {
		var strPlanned sql.NullString
		var strFact sql.NullString
		item := new(model.CarriageRoute)

		if err := rows.Scan(&item.Id, &item.CarriageId, &item.CheckPointId, &strPlanned, &strFact); err != nil {
			return nil, err
		}

		if item.Planned, err = str2DateRFC3339(strPlanned); err != nil {
			return nil, err
		}
		if item.Fact, err = str2DateRFC3339(strFact); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
