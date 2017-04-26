package psql

import (
	"TruckMonitor-Backend/dao"
	"TruckMonitor-Backend/model"
	"database/sql"
)

type psqlCheckPoint struct {
	context PsqlContext
}

func CheckPointDao(context PsqlContext) dao.CheckPointDao {
	return &psqlCheckPoint{context}
}

func (dao *psqlCheckPoint) db() *sql.DB {
	return dao.context.GetDb()
}

func (dao *psqlCheckPoint) FindById(id int) (*model.CheckPoint, error) {
	var data model.CheckPoint
	row := dao.db().QueryRow("SELECT * FROM check_point WHERE id=$1", id)
	err := row.Scan(&data.Id, &data.Name, &data.Address, &data.Latitude, &data.Longitude)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
