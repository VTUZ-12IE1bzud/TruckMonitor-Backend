package model

type (
	// Вид ТС
	VehicleType struct {
		Id   int
		Name string
	}

	// Марка ТС
	VehicleBrand struct {
		Id     int
		TypeId int
		Name   string
	}

	// Марка ТС
	VehicleModel struct {
		Id          int
		BrandId     int
		Name        string
		CapacityMax float32
	}

	// ТС
	Vehicle struct {
		Id           int
		ModelId      int
		Vin          string
		LicencePlate float32
		RegionNumber int
	}
)
