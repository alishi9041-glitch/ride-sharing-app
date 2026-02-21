package models

type Vehicle struct {
	VehicleID string      `json:"vehicleId"`
	Type      VehicleType `json:"vehicleType"`
}
