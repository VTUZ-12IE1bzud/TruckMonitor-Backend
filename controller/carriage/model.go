package carriage

import "time"

type (

	currentResponse struct {
		Id         int           `json:"id"`
		Carriages  []carriage    `json:"carriages"`
		CheckPoint []checkPoint `json:"checkPoints"`
	}

	carriage struct {
		Contract      contract       `json:"contract"`
		Nomenclatures []nomenclature `json:"nomenclatures"`
	}

	contract struct {
		Number      string   `json:"number"`
		CustomsLink string   `json:"customsLink"`
		From        store    `json:"from"`
		Before      store    `json:"before"`
		Packaging   string   `json:"packaging"`
		Manager     manager `json:"manager"`
	}

	store struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	}

	nomenclature struct {
		Name    string  `json:"name"`
		Amount  float32 `json:"amount"`
		Measure string  `json:"measure"`
	}

	manager struct {
		Name      string `json:"name"`
		PhotoLink string `json:"photolink"`
		Phone     string `json:"phone"`
	}

	checkPoint struct {
		Coordinates coordinates `json:"coordinates"`
		Planned     time.Time       `json:"planned"`
		Fact        *time.Time       `json:"fact"`
	}

	coordinates struct {
		Latitude  float32 `json:"latitude"`
		Longitude float32 `json:"longitude"`
	}
)
