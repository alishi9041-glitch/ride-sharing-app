package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type InventoryHandler struct {
}

func NewInventoryHandler() InventoryHandler {
	return InventoryHandler{}
}

func (ih *InventoryHandler) InventoryHealth(c echo.Context) error {
	storeID := c.Param("storeId")
	if storeID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Store ID is required",
		})
	}

	return c.JSON(http.StatusOK, "Ok")
}
