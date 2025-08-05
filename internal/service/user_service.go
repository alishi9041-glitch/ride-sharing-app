package service

import (
	"fmt"

	"joi-delivery-golang/internal/models"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) GetCartByUserID(userId string) (*models.Cart, error) {
	cart, exists := usersCart[userId]
	if !exists {
		return nil, fmt.Errorf("cart not found")
	}

	return cart, nil
}
