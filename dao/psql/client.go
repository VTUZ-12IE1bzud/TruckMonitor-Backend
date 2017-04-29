package psql

import (
	"TruckMonitor-Backend/dao"
	"TruckMonitor-Backend/model"
	"database/sql"
)

type psqlClient struct {
	context PsqlContext
}

func ClientDao(context PsqlContext) dao.ClientDao {
	return &psqlClient{context}
}

func (dao *psqlClient) db() *sql.DB {
	return dao.context.GetDb()
}

func (dao *psqlClient) FindById(id int) (*model.Client, error) {
	var data model.Client
	row := dao.db().QueryRow("SELECT * FROM client WHERE id=$1", id)
	if err := row.Scan(&data.Id, &data.Name, &data.Itn, &data.Iec, &data.Address); err != nil {
		return nil, err
	}
	return &data, nil
}
