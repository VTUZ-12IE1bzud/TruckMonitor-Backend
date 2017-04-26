package model

import "time"

type (
	// Договор
	Contract struct {
		Id                      int
		ManagerId               int
		ClientId                int
		PackagingId             int
		StoreFromId             int
		StoreBeforeId           int
		Number                  string
		Price                   float32
		ConfirmationPaymentLink string
		ConfirmationCustomsLink string
		DateShipment            time.Time
	}

	// Предмет договора
	ContractDetail struct {
		Id             int
		ContractId     int
		NomenclatureId int
		Amount         float32
		Price          float32
	}
)
