package service

import (
	"basic/internal/models"
	"math"
	"sort"
)

const MIN_DISTANCE = 5

type MatchDriver interface {
	findDriversNearby(riderLoc models.GeoLocation) []*models.Driver
	notifyDrivers(rideID string, drivers []*models.Driver)
}

type Match struct{}

type DriverInfo struct {
	driver   *models.Driver
	distance float64
	eta      int
}

// TODO: maybe increasing the distance from x to x+10kms after every 2 mins of no acceptance?
func (m *Match) findDriversNearby(riderLoc models.GeoLocation, vehicleType models.VehicleType) []*DriverInfo {

	// filter all drivers within x kms of RidersLoc and are of available location
	var nearbyDrivers []*DriverInfo
	for _, driver := range drivers {

		if driver.Status == models.AVAILABLE && driver.Vehicle.Type == vehicleType {
			// calculate distance
			dist := calculateDistance(driver.GeoLocation, riderLoc)
			if dist <= MIN_DISTANCE {
				// nearbyDrivers = append(nearbyDrivers, v)
				eta := int(math.Ceil(calculateETA(dist, 30)))
				nearbyDrivers = append(nearbyDrivers, &DriverInfo{
					driver:   driver,
					distance: dist,
					eta:      eta,
				})
			}
		}
	}

	// sort on the basis of priority
	sort.Slice(nearbyDrivers, func(i, j int) bool {

		if nearbyDrivers[i].distance != nearbyDrivers[j].distance {
			return nearbyDrivers[i].distance < nearbyDrivers[j].distance
		}

		if nearbyDrivers[i].driver.Rating != nearbyDrivers[j].driver.Rating {
			return nearbyDrivers[i].driver.Rating > nearbyDrivers[j].driver.Rating
		}

		return nearbyDrivers[i].eta < nearbyDrivers[j].eta
	})

	return nearbyDrivers
}

func (m *Match) notifyDrivers(rideID string, drivers []*DriverInfo) {
	for _, driverInfo := range drivers {
		driverInfo.driver.NotifyDriver(rideID)
	}
}

func calculateETA(distanceKm float64, avgSpeedKmph float64) float64 {
	if distanceKm <= 0 || avgSpeedKmph <= 0 {
		return 0
	}

	return (distanceKm / avgSpeedKmph) * 60
}
