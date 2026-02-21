package response

import "basic/internal/models"

type BookARideResponse struct {
	RideID   string  `json:"rideID"`
	DriverID string  `json:"driverID"`
	ETA      int     `json:"eta"` //in minutes
	Fare     float64 `json:"fare"`
}

type AcceptRideResponse struct {
	DriverID string `json:"driverID"`
}

type StartRideResponse struct {
	RideID string `json:"rideID"`
}

type CompleteRideResponse struct {
	RideID        string                `json:"rideID"`
	Status        models.RIDE_STATUS    `json:"status"`
	FinalFare     float64               `json:"finalFare"`
	PaymentStatus models.PAYMENT_STATUS `json"paymentStatus`
}
