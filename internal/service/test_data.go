package service

import (
	"joi-delivery-golang/internal/models"
)

var (
	stores    = make(map[string]*models.Store)
	products  = make(map[string]*models.GroceryProduct)
	usersCart = make(map[string]*models.Cart)
	users     = make(map[string]*models.User)
)

func InitializeTestData() {
	users = map[string]*models.User{
		"user101": {
			ID:        "user101",
			FirstName: "John",
			LastName:  "Doe",
		},
		"user102": {
			ID:        "user102",
			FirstName: "Jane",
			LastName:  "Smith",
		},
	}

	stores = map[string]*models.Store{
		"store101": {
			Outlet: models.Outlet{
				ID:   "store101",
				Name: "Fresh Picks",
			},
		},
		"store102": {
			Outlet: models.Outlet{
				ID:   "store102",
				Name: "Natural Choice",
			},
		},
	}

	products = map[string]*models.GroceryProduct{
		"product101": newProduct("product101", "Wheat Bread", "store101"),
		"product102": newProduct("product102", "Spinach", "store102"),
		"product103": newProduct("product103", "Crackers", "store101"),
	}

	usersCart = map[string]*models.Cart{
		"user101": newCart("cart101", "user101"),
		"user102": newCart("cart102", "user102"),
	}
}

func newCart(cartId, userId string) *models.Cart {
	return &models.Cart{
		ID:       cartId,
		User:     users[userId],
		Products: make([]models.GroceryProduct, 0),
	}
}

func newProduct(productId, name, storeId string) *models.GroceryProduct {
	return &models.GroceryProduct{
		Product: models.Product{
			ID:   productId,
			Name: name,
			MRP:  10.5,
		},
		Threshold:      10,
		AvailableStock: 30,
		Weight:         500.00,
		Store:          *stores[storeId],
	}

}

func ClearTestData() {
	stores = make(map[string]*models.Store)
	products = make(map[string]*models.GroceryProduct)
	users = make(map[string]*models.User)
	usersCart = make(map[string]*models.Cart)
}
