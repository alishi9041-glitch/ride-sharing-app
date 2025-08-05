package models

type Cart struct {
	ID       string           `json:"id"`
	Outlet   *Outlet          `json:"outlet"`
	User     *User            `json:"user"`
	Products []GroceryProduct `json:"products"`
}
