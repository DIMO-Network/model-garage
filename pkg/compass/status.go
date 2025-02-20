package compass

// ConvertPSIToKPa converts a pressure value from psi to kPa.
func ConvertPSIToKPa(psi float64) float64 {
	return psi * 6.89476
}

// ConvertBarToKPa converts a pressure value from bar to kPa.
func ConvertBarToKPa(bar float64) float64 {
	return bar * 100
}
