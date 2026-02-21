package models

type Rider struct {
	ID          string      `json:"id"`
	Username    string      `json:"username"`
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	GeoLocation GeoLocation `json:"geolocation"`
}

type GeoLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
