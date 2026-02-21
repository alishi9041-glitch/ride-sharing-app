package service

import (
	"basic/internal/models"
	"errors"
	"fmt"
	"sync"
	"time"

	"basic/internal/fare"

	"github.com/google/uuid"
	"github.com/umahmood/haversine"
)

var rides = make(map[string]*Ride)

type Ride struct {
	ID        string
	RiderID   string
	DriverID  string
	Pickup    models.GeoLocation
	Drop      models.GeoLocation
	Status    models.RIDE_STATUS
	Fare      float64
	CreatedAt time.Time
	UpdatedAt time.Time
	ETA       int
	mu        sync.Mutex
}

type RideService struct {
	fareCalculator fare.FareCalculator
}

func NewRideService(fareCalculator fare.FareCalculator) *RideService {
	return &RideService{
		fareCalculator: fareCalculator,
	}
}

type RideServiceI interface {
	RequestRide(riderID string, pickup, drop models.GeoLocation, vehicleType models.VehicleType) (*Ride, error)
	AcceptRide(rideID string, driverID string) (*Ride, error)
	StartRide(rideID string) (*Ride, error)
	CompleteRide(rideID string) (*Ride, error)
	CancelRide(rideID string) (*Ride, error)
}

func (rs *RideService) RequestRide(riderId string, pickup, drop models.GeoLocation, vehicleType models.VehicleType) (*Ride, error) {
	fmt.Printf("Reaching here")

	//Validate rider
	if _, ok := riders[riderId]; !ok {
		// log.Fatalf("Invalid rider passed.")
		return nil, errors.New("Rider does not exist.")
	}

	// validate drop and pickup not same, should be withing city, etc.
	// TODO: think of more validations, maybe add validations in a different method or file
	if pickup.Latitude == drop.Latitude && pickup.Longitude == drop.Longitude {
		// log.Fatalf("Pick up and drop location same.")
		fmt.Printf("Pick up and drop location same.")
		return nil, errors.New("Please choose a destination different from current location.")
	}

	// calculate Fare for the distance
	distance := calculateDistance(pickup, drop)

	if distance < 0 {
		return nil, errors.New("Please choose a different destination.")
	}

	fare := rs.fareCalculator.Calculate(fare.PricingInput{
		DistanceKm:     distance,
		RequestTime:    time.Now(),
		DiscountAmount: 0,
		SurgePrice:     4, // TODO: revisit
	})

	ride := &Ride{
		ID:        uuid.New().String(),
		RiderID:   riderId,
		Pickup:    pickup,
		Drop:      drop,
		Status:    models.REQUESTED,
		Fare:      fare,
		CreatedAt: time.Now(),
	}

	rides[ride.ID] = ride
	// start matching algorithm, notify drivers
	match := &Match{}

	drivers := match.findDriversNearby(pickup, vehicleType)

	if len(drivers) == 0 {
		return nil, errors.New("No driver found currently")
	}

	ride.ETA = drivers[0].eta
	// notify the drivers
	match.notifyDrivers(ride.ID, drivers)

	// create ride id and return
	fmt.Println("Rides created %v", rides)
	return ride, nil

}

func (rs *RideService) AcceptRide(rideID string, driverID string) error {

	// get ride details and check if its still pending
	ride, ok := rides[rideID]
	if !ok {
		fmt.Println("Ride ID incorrect")
		return errors.New("Ride not found")
	}
	// lock ride so that no other driver can take action on it
	ride.mu.Lock()
	defer ride.mu.Unlock()

	if ride.Status != models.REQUESTED {
		//Already accepted
		fmt.Println("Ride no longer avaulable")
		return errors.New("Ride is no longer available")
	}

	// check driver is in available state
	driver, ok := drivers[driverID]
	// this scenario should never happen
	if !ok {
		fmt.Println("Driver not found")
		return errors.New("Driver not found")
	}

	if driver.Status != models.AVAILABLE {
		if !ok {
			if !ok {
				fmt.Println("You have an ongoing ride already. Please complete it to accept a new ride.")
				return errors.New("You have an ongoing ride already. Please complete it to accept a new ride.")
			}
		}
	}

	rideFound := false

	for _, request := range driver.RideRequests {
		if rideID == request {
			rideFound = true
		}
	}

	if !rideFound {
		return errors.New("Incorrect ride!")
	}

	// update ride and driver status
	driver.Status = models.BUSY
	driver.RideRequests = removePendingRide(driver.RideRequests, rideID)
	ride.Status = models.ACCEPTED
	ride.DriverID = driverID

	return nil
}

func (rs *RideService) StartRide(rideID string) (*Ride, error) {

	ride, ok := rides[rideID]
	if !ok {
		return nil, errors.New("Ride not found")
	}

	if ride.Status != models.ACCEPTED {
		return nil, errors.New("Ride is not accepted")
	}

	// update ride status
	ride.Status = models.STARTED
	return ride, nil
}

func (rs *RideService) CompleteRide(rideID string) (*Ride, error) {

	ride, ok := rides[rideID]
	if !ok {
		return nil, errors.New("Ride not found")
	}

	if ride.Status != models.STARTED {
		return nil, errors.New("Ride is not started yet")
	}

	// update ride status
	ride.Status = models.COMPLETED

	// Free the driver
	driver := drivers[ride.DriverID]
	driver.Status = models.AVAILABLE

	return ride, nil

}

func (rs *RideService) CancelRide(rideID string) (*Ride, error) {

	ride, ok := rides[rideID]
	if !ok {
		return nil, errors.New("Ride not found")
	}

	if ride.Status == models.ACCEPTED {
		return nil, errors.New("Can't cancel this ride")
	}

	// update ride status
	ride.Status = models.CANCELLED

	// Free the driver
	driver := drivers[ride.DriverID]
	driver.Status = models.AVAILABLE

	return ride, nil

}

// should go to some utility package maybe
func calculateDistance(source models.GeoLocation, dest models.GeoLocation) float64 {
	loc1 := haversine.Coord{Lat: source.Latitude, Lon: source.Longitude}
	loc2 := haversine.Coord{Lat: dest.Latitude, Lon: dest.Longitude}

	_, km := haversine.Distance(loc1, loc2)

	return km
}

func removePendingRide(slice []string, value string) []string {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
