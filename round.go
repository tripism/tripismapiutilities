package tripismapiutilities

import (
	"fmt"
	"strconv"
)

// RoundFloat rounds float to provided precision
func RoundFloat(pattern string, f float64) (rounded float64) {
	if pattern == "" {
		return
	}
	rounded, err := strconv.ParseFloat(fmt.Sprintf(pattern, f), 64)
	if err != nil {
		return
	}

	return
}
