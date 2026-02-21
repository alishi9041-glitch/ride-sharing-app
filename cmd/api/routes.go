package api

import (
	"basic/cmd/api/handler"

	"github.com/labstack/echo/v4"
)

func BindRoutes(server *echo.Echo, hand handler.Handler) {
	//Ride routes
	server.POST("/book-ride", hand.BookARide)
	server.POST("/ride/accept", hand.AcceptRide)
	server.POST("/ride/start", hand.StartRide)
	server.POST("/ride/complete", hand.CompleteRide)
}
