package psql

import (
	"TruckMonitor-Backend/dao"
	"database/sql"
	"TruckMonitor-Backend/model"
)

type psqlStore struct {
	context PsqlContext
}

func StoreDao(context PsqlContext) dao.StoreDao {
	return &psqlStore{context}
}

func (dao *psqlStore) db() *sql.DB {
	return dao.context.GetDb()
}

func (dao *psqlStore) FindById(id int) (*model.Store, error) {
	var data model.Store
	row := dao.db().QueryRow("SELECT * FROM store WHERE id=$1", id)
	err := row.Scan(&data.Id, &data.Name, &data.Address, &data.Latitude, &data.Longitude)
	if err != nil {
		return nil, err
	}
	return &data, nil
}