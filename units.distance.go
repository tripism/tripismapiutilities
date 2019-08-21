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

// MilesToMetersIntFloat converts distnace in miles to meters in float64
func MilesToMetersIntFloat(miles int) float64 {
	return float64(miles) * 1609.344
}

// MilesToMetersFloatFloat converts distance in miles to meters in float64
func MilesToMetersFloatFloat(miles float64) float64 {
	return miles * 1609.344
}

// MetersToMiles converts distance in meters to miles
func MetersToMiles(meters int) int {
	return MetersToMilesIntInt(meters)
}

// MetersToMilesIntInt converts distance in meters to miles in integer
func MetersToMilesIntInt(meters int) int {
	miles := float64(meters) / 1609.344
	s := fmt.Sprintf("%.0f", miles)
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	t := int(miles)
	return t
}

// MetersToMilesIntFloat converts distance in meters to miles in float64
func MetersToMilesIntFloat(meters int) float64 {
	return float64(meters) / 1609.344
}

// MetersToMilesFloatFloat converts distance in meters to miles in float64
func MetersToMilesFloatFloat(meters float64) float64 {
	return meters / 1609.344
}
