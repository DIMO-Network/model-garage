package teslatelem

const kilometersPerMile = 1.609344

func barsToKilopascals(bars float64) float64 {
	return 100 * bars
}

func milesToKilometers(miles float64) float64 {
	return kilometersPerMile * miles
}

func kilowattsToWatts(kilowatts float64) float64 {
	return 1000 * kilowatts
}
