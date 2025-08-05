package service

import "joi-delivery-golang/internal/models"

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (ps *ProductService) GetProductByID(productId, outletId string) *models.GroceryProduct {
	product, _ := products[productId]
	if product.Store.ID != outletId {
		return nil
	}

	return product
}
