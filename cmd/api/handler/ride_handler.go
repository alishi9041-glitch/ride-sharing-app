package handler

import (
	"basic/internal/dto/request"
	"basic/internal/dto/response"
	"basic/internal/models"
	"basic/internal/service"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RideHandler struct {
	rideService *service.RideService
}

func NewRideHandler(rideSvc *service.RideService) RideHandler {
	return RideHandler{
		rideService: rideSvc,
	}
}

func (rh *RideHandler) BookARide(c echo.Context) error {
	var req request.BookARideRequest

	// bind the type
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Success": false,
			"Message": "Invalid request format",
		})
	}

	// validate required params
	if req.UserID == "" || req.DestPos.Latitude == 0 || req.DestPos.Longitude == 0 || req.CurrentPos.Latitude == 0 || req.CurrentPos.Longitude == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Success": false,
			"Message": "Missing required fields",
		})
	}

	ride, err := rh.rideService.RequestRide(req.UserID, req.CurrentPos, req.DestPos, req.VehicleType)

	if err != nil {
		// log.Fatal("Error while requesting a ride ", err)
		fmt.Printf("Error while requesting a ride %v", err)

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Success": false,
			"Message": "Internal server error",
		})
	}

	return c.JSON(http.StatusOK, ride)

}

func (rh *RideHandler) AcceptRide(c echo.Context) error {
	var req request.AcceptRideRequest

	// bind the type
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Success": false,
			"Message": "Invalid request format",
		})
	}

	// validate required params
	if req.RideID == "" || req.DriverID == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Success": false,
			"Message": "Missing required fields",
		})
	}

	err := rh.rideService.AcceptRide(req.RideID, req.DriverID)

	if err != nil {
		// log.Fatal("Error while accepting the ride ", err)

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Success": false,
			"Message": "Internal server error",
		})
	}

	return nil
}

func (rh *RideHandler) StartRide(c echo.Context) error {
	var req request.StartRideRequest

	// bind the type
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Success": false,
			"Message": "Invalid request format",
		})
	}

	// validate required params
	if req.RideID == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Success": false,
			"Message": "Missing required fields",
		})
	}

	ride, err := rh.rideService.StartRide(req.RideID)

	if err != nil {
		// log.Fatal("Error while accepting the ride ", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Success": false,
			"Message": "Internal server error",
		})
	}

	return c.JSON(http.StatusOK, ride.ID)
}

func (rh *RideHandler) CompleteRide(c echo.Context) error {
	var req request.CompleteRideRequest

	// bind the type
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Success": false,
			"Message": "Invalid request format",
		})
	}

	// validate required params
	if req.RideID == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Success": false,
			"Message": "Missing required fields",
		})
	}

	ride, err := rh.rideService.CompleteRide(req.RideID)

	if err != nil {
		// log.Fatal("Error while completing the ride ", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Success": false,
			"Message": "Internal server error",
		})
	}

	return c.JSON(http.StatusOK, &response.CompleteRideResponse{
		RideID:        ride.ID,
		Status:        models.COMPLETED,
		FinalFare:     ride.Fare,
		PaymentStatus: models.PENDING,
	})
}
