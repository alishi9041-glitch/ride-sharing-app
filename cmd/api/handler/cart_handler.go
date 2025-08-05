package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"joi-delivery-golang/internal/dto/request"
	"joi-delivery-golang/internal/service"
)

type CartHandler struct {
	cartService *service.CartService
}

func NewCartHandler(cartService *service.CartService) CartHandler {
	return CartHandler{
		cartService: cartService,
	}
}

func (ch *CartHandler) AddToCart(c echo.Context) error {
	var req request.AddToCartRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Success": false,
			"Message": "Invalid request format",
		})
	}

	if req.UserID == "" || req.OutletID == "" || req.ProductID == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Success": false,
			"Message": "Missing required fields",
		})
	}

	resp, err := ch.cartService.AddToCart(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Success": false,
			"Message": "Internal server error",
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (ch *CartHandler) GetCart(c echo.Context) error {
	userId := c.QueryParam("userId")
	if userId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "User ID is required",
		})
	}

	cart, err := ch.cartService.GetCartByUserID(userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, cart)
}
