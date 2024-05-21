// Package weather provides tools to consult the current weather conditions.
package weather

// CurrentCondition is the current weather conditon.
var CurrentCondition string

// CurrentLocation is the current location.
var CurrentLocation string

// Forecast returns a string with the current location and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
