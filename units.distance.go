package tripismapiutilities

import (
	"fmt"
	"strconv"
)

// MilesToMeters converts distance in miles to meters
func MilesToMeters(miles int) int {
	return MilesToMetersIntInt(miles)
}

// MilesToMetersIntInt converts distance in miles to meters in integer
func MilesToMetersIntInt(miles int) int {
	meters := float64(miles) * 1609.344
	s := fmt.Sprintf("%.0f", meters)
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	t := int(meters)
	return t
}

func MilesToMeters_IntFloat(miles int) float64 {
	return float64(miles) * 1609.344
}

func MilesToMeters_FloatFloat(miles float64) float64 {
	return miles * 1609.344
}

func MetersToMiles(meters int) int {
	return MetersToMiles_IntInt(meters)
}

func MetersToMiles_IntInt(meters int) int {
	miles := float64(meters) / 1609.344
	s := fmt.Sprintf("%.0f", miles)
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	t := int(miles)
	return t
}

func MetersToMiles_IntFloat(meters int) float64 {
	return float64(meters) / 1609.344
}

func MetersToMiles_FloatFloat(meters float64) float64 {
	return meters / 1609.344
}
