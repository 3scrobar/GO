// Package weather : For Weather Forecast, contains Forecast.
package weather

var (
	//CurrentCondition : Conditions for the localisation.
	CurrentCondition string
	//CurrentLocation : Current location.
	CurrentLocation string
)

// Forecast : Give the current forecast. Need Location , Conditions.Return a string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
