package request

import "basic/internal/models"

type AcceptRideRequest struct {
	RideID   string `json:"rideID"`
	DriverID string `json:"driverID"`
}

type BookARideRequest struct {
	UserID      string             `json:"userId"`
	CurrentPos  models.GeoLocation `json:"currentPosition"`
	DestPos     models.GeoLocation `json:"destinationPosition"`
	VehicleType models.VehicleType `json:"vehicleType"`
}

type StartRideRequest struct {
	RideID string `json:"rideID"`
}

type CompleteRideRequest struct {
	RideID string `json:"rideID"`
}
