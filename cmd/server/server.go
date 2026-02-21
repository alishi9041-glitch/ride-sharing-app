package server

import (
	"context"

	"basic/cmd/api"
	"basic/cmd/api/handler"
	"basic/internal/fare"
	"basic/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	BASE_FARE   = 50
	RATE_PER_KM = 10
	NIGHT_PRICE = 1.2
)

type Server struct {
	server   *echo.Echo
	handlers handler.Handler
}

func NewServer(ctx context.Context) *Server {
	server := echo.New()
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.CORS())

	s := &Server{
		server:   server,
		handlers: registerHandlers(),
	}

	api.BindRoutes(server, s.handlers)

	service.InitializeTestData()

	return s
}

func (s *Server) Start(port string) error {
	return s.server.Start(port)
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func registerHandlers() handler.Handler {

	calculator := fare.NewFareCalculator([]fare.FareStrategy{
		fare.BaseFareStrategy{BaseFare: float64(BASE_FARE)},
		fare.PerKmStrategy{RatePerKm: float64(RATE_PER_KM)},
		fare.SurgeStrategy{},
		fare.NightStrategy{Multiplier: NIGHT_PRICE},
		fare.DiscountStrategy{},
	})

	rideService := service.NewRideService(calculator)
	rideHandler := handler.NewRideHandler((rideService))

	return handler.Handler{
		RideHandler: rideHandler,
	}
}
