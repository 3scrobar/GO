package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100.0
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	mult := float64(productionRate) * successRate / 100.0
	mult /= 60
	return int(mult)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	cost10 := 95000
	cost := 10000
	var total uint
	var unit int
	if carsCount > 9 {
		dec := carsCount / 10
		unit = carsCount % 10
		total = uint(dec*cost10 + unit*cost)
	} else {
		total = uint(carsCount * cost)
	}
	return total
}
