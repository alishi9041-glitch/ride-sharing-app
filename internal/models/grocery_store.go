package models

type Store struct {
	Outlet
	Inventory map[string]interface{} `json:"inventory"`
}
