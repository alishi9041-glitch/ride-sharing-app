package models

type User struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Cart        *Cart  `json:"cart"`
}
