package tripismapiutilities

// CelciusToFahrenheit convert temperature from Celcius to Fahrenheit
// https://www.rapidtables.com/convert/temperature/how-celsius-to-fahrenheit.html
func CelciusToFahrenheit(celcius float64) float64 {
	return celcius*9/5 + 32
}

// MetersPerSecondToMilesPerHour convert wind speed from meters per second to miles per hour
// https://www.weather.gov/media/epz/wxcalc/windConversion.pdf
func MetersPerSecondToMilesPerHour(mps float64) float64 {
	return mps * 2.23694
}
