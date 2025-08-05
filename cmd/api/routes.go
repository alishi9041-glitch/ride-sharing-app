package api

import (
	"github.com/labstack/echo/v4"
	"joi-delivery-golang/cmd/api/handler"
)

func BindRoutes(server *echo.Echo, hand handler.Handler) {
	// Cart routes
	server.GET("/cart/view", hand.GetCart)
	server.POST("/cart/product", hand.AddToCart)

	// Inventory routes
	server.POST("/inventory/health", hand.InventoryHealth)
}
