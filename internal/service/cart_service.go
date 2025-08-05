package service

import (
	"joi-delivery-golang/internal/dto/request"
	"joi-delivery-golang/internal/dto/response"
	"joi-delivery-golang/internal/models"
)

type CartService struct {
	userSvc    *UserService
	productSvc *ProductService
}

func NewCartService(userSvc *UserService, productSvc *ProductService) *CartService {
	return &CartService{
		userSvc:    userSvc,
		productSvc: productSvc,
	}
}

func (cs *CartService) AddToCart(req request.AddToCartRequest) (*response.AddToCartResponse, error) {
	product := cs.getProductByID(req.ProductID, req.OutletID)

	cart := cs.findCartByUser(req.UserID)

	cart.Products = append(cart.Products, *product)

	return &response.AddToCartResponse{
		Cart: *cart,
	}, nil
}

func (cs *CartService) getProductByID(productId, outletId string) *models.GroceryProduct {
	return cs.productSvc.GetProductByID(productId, outletId)
}

func (cs *CartService) findCartByUser(userId string) *models.Cart {
	return usersCart[userId]
}

func (cs *CartService) GetCartByUserID(userId string) (*models.Cart, error) {
	return cs.userSvc.GetCartByUserID(userId)
}
