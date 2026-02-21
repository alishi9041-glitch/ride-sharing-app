package service

import (
	"basic/internal/models"
)

var (
	riders   = make(map[string]*models.Rider)
	drivers  = make(map[string]*models.Driver)
	vehicles = make(map[string]*models.Vehicle)
)

func InitializeTestData() {

	riders = map[string]*models.Rider{
		"user101": {
			ID:        "user101",
			FirstName: "John",
			LastName:  "Doe",
			Username:  "username101",
			GeoLocation: models.GeoLocation{
				Latitude:  1.1,
				Longitude: 5.22,
			},
		},
		"user102": {
			ID:        "user102",
			FirstName: "Jane",
			LastName:  "Smith",
			Username:  "username102",
			GeoLocation: models.GeoLocation{
				Latitude:  7.1,
				Longitude: 1.22,
			},
		},
	}

	drivers = map[string]*models.Driver{
		"driver101": {
			ID:        "driver101",
			FirstName: "Jai",
			LastName:  "Vaswani",
			GeoLocation: models.GeoLocation{
				Latitude:  18.52,
				Longitude: 73.85,
			},
			Status:  models.AVAILABLE,
			Vehicle: *newVehicle("vehicle101", models.SUV),
		},
		"driver102": {
			ID:        "driver102",
			FirstName: "Robert",
			LastName:  "Smith",
			GeoLocation: models.GeoLocation{
				Latitude:  1.1,
				Longitude: 5.22,
			},
			Status:  models.AVAILABLE,
			Vehicle: *newVehicle("vehicle102", models.Sedan),
		},
	}

	// vehicles = map[string]*models.Vehicle{
	// 	"vehicle01": newVehicle("vehicle101", models.SUV),
	// 	"vehicle02": newVehicle("vehicle102", models.Sedan),
	// }

}

func newVehicle(vehicleId string, Type models.VehicleType) *models.Vehicle {
	return &models.Vehicle{
		VehicleID: vehicleId,
		Type:      Type,
	}

}

func ClearTestData() {
	vehicles = make(map[string]*models.Vehicle)
	riders = make(map[string]*models.Rider)
	drivers = make(map[string]*models.Driver)
}
