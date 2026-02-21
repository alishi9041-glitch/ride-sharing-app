package models

type Driver struct {
	ID           string        `json:"id"`
	FirstName    string        `json:"firstName"`
	LastName     string        `json:"lastName"`
	Status       DRIVER_STATUS `json:"status"`
	GeoLocation  GeoLocation   `json:"geolocation"`
	RideRequests []string      `json:"rideRequests"`
	Vehicle      Vehicle       `json:"vehicle"`
	Rating       int           `json:"rating"`
}

type DriverLogic interface {
	notifyDriver(rideID string) bool
}

func (d *Driver) NotifyDriver(rideID string) {
	if d.Status == AVAILABLE {
		d.RideRequests = append(d.RideRequests, rideID)
	}
}
