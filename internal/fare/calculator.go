package fare

import "time"

type PricingInput struct {
	DistanceKm     float64
	SurgePrice     float64
	RequestTime    time.Time
	DiscountAmount float64
}

type FareStrategy interface {
	Calculate(currentFare float64, priceInput PricingInput) float64
}

type FareCalculator interface {
	Calculate(priceInput PricingInput) float64
}

type fareCalculator struct {
	strategies []FareStrategy
}

func NewFareCalculator(strategies []FareStrategy) FareCalculator {
	return &fareCalculator{
		strategies: strategies,
	}
}

func (f *fareCalculator) Calculate(priceInput PricingInput) float64 {

	fare := 0.0

	for _, strategy := range f.strategies {
		fare = strategy.Calculate(fare, priceInput)
	}

	if fare < 0 {
		return 0
	}

	return fare
}
