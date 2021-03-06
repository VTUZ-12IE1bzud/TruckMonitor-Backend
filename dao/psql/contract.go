package psql

import (
	"TruckMonitor-Backend/dao"
	"TruckMonitor-Backend/model"
	"database/sql"
)

type psqlContract struct {
	context PsqlContext
}

func ContractDao(context PsqlContext) dao.ContractDao {
	return &psqlContract{context}
}

func (dao *psqlContract) db() *sql.DB {
	return dao.context.GetDb()
}

func (dao *psqlContract) FindById(id int) (*model.Contract, error) {
	var data model.Contract
	row := dao.db().QueryRow("SELECT * FROM contract WHERE id=$1", id)
	if err := row.Scan(&data.Id, &data.ManagerId, &data.ClientId, &data.PackagingId, &data.StoreFromId,
		&data.StoreBeforeId, &data.Number, &data.Price, &data.ConfirmationPaymentLink,
		&data.ConfirmationCustomsLink, &data.DateShipment); err != nil {
		return nil, err
	}
	return &data, nil
}

func (dao *psqlContract) FindDetails(contractId int) ([]*model.ContractDetail, error) {
	rows, err := dao.db().Query("SELECT * FROM contract_detail WHERE contract_id=$1", contractId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := make([]*model.ContractDetail, 0)
	for rows.Next() {
		item := new(model.ContractDetail)
		if err := rows.Scan(&item.Id, &item.ContractId, &item.NomenclatureId, &item.Amount, &item.Price); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
