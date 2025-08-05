package service

import (
	"joi-delivery-golang/internal/dto/request"
	"testing"

	"github.com/stretchr/testify/assert"
	"joi-delivery-golang/internal/models"
)

func setupTestData() {
	stores = make(map[string]*models.Store)
	products = make(map[string]*models.GroceryProduct)
	usersCart = make(map[string]*models.Cart)
	users = make(map[string]*models.User)

	InitializeTestData()
}

func TestCartService_AddToCart_Success_Store(t *testing.T) {
	setupTestData()
	userSvc := NewUserService()
	productSvc := NewProductService()
	cs := NewCartService(userSvc, productSvc)

	req := request.AddToCartRequest{
		UserID:    "user101",
		OutletID:  "store101",
		ProductID: "product101",
	}

	response, err := cs.AddToCart(req)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Cart)
	assert.NotEmpty(t, response.Cart.ID)
}

func TestCartService_GetCart_Success(t *testing.T) {
	setupTestData()
	userSvc := NewUserService()
	productSvc := NewProductService()
	cs := NewCartService(userSvc, productSvc)

	req := request.AddToCartRequest{
		UserID:    "user101",
		OutletID:  "store101",
		ProductID: "product101",
	}

	response, err := cs.AddToCart(req)
	assert.NoError(t, err)

	cart, err := cs.GetCartByUserID(response.Cart.User.ID)
	assert.NoError(t, err)
	assert.NotNil(t, cart)
	assert.NotEmpty(t, cart)
}
