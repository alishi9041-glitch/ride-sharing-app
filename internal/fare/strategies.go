package fare

type BaseFareStrategy struct {
	BaseFare float64
}

func (b BaseFareStrategy) Calculate(currentFare float64, priceInput PricingInput) float64 {
	return currentFare + b.BaseFare
}

type PerKmStrategy struct {
	RatePerKm float64
}

func (p PerKmStrategy) Calculate(currentFare float64, priceInput PricingInput) float64 {
	return currentFare + (priceInput.DistanceKm * p.RatePerKm)
}

type SurgeStrategy struct {
}

func (s SurgeStrategy) Calculate(currentFare float64, priceInput PricingInput) float64 {
	if priceInput.SurgePrice > 1 {
		return currentFare * priceInput.SurgePrice
	}
	return currentFare
}

type NightStrategy struct {
	Multiplier float64
}

func (n NightStrategy) Calculate(currentFare float64, priceInput PricingInput) float64 {
	hour := priceInput.RequestTime.Hour()

	// considering night
	if hour >= 22 || hour < 6 {
		return currentFare * n.Multiplier
	}

	return currentFare
}

type DiscountStrategy struct{}

func (d DiscountStrategy) Calculate(currentFare float64, priceInput PricingInput) float64 {
	return currentFare - priceInput.DiscountAmount
}
