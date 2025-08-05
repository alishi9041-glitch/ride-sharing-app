package response

import "joi-delivery-golang/internal/models"

type AddToCartResponse struct {
	Cart models.Cart `json:"cart"`
}
