package models

type Product struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	MRP  float64 `json:"mrp"`
}
