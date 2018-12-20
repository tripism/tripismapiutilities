package tripismapiutilities

import (
	"runtime"
)

// getMaxProcs returns the maximum processor count allocation for running GoRoutines
func GetMaxProcs(requestedValue int) int {
	currentNumProc := runtime.NumCPU()

	// If we only have 2 or less  processors available, allocate maximum 1 to the GOMAXPROCS
	if currentNumProc < 3 {
		return 1
	}

	//Always leave one processor core free
	currentNumProc = currentNumProc - 1

	if requestedValue > currentNumProc {
		return currentNumProc
	}

	return requestedValue
}
