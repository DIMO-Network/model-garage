// Package unit converts between commonly used units of measure for
// vehicle signals.
package unit

func KilowattsToWatts(x float64) float64 {
	return 1000 * x
}

const kilopascalsPerAtmosphere = 101.325

func AtmospheresToKilopascals(x float64) float64 {
	return kilopascalsPerAtmosphere * x
}

const kilometersPerMile = 1.609344

func MilesToKilometers(x float64) float64 {
	return kilometersPerMile * x
}

func MilesPerHourToKilometersPerHour(x float64) float64 {
	return kilometersPerMile * x
}
