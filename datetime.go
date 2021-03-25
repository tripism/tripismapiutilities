package tripismapiutilities

import (
	"strings"
	"time"
)

// MinutesToHoursAndMinutes formats a number of minutes into a time duration representation in the format XhXm with optional override of the h and m text
func MinutesToHoursAndMinutes(d int, hourFormat string, minuteFormat string) string {
	td := time.Minute * time.Duration(d)
	// String returns a string representing the duration in the form "72h3m0.5s".
	// Leading zero units are omitted.
	tds := td.String()
	// remove the seconds representation
	tda := strings.TrimSuffix(tds, "0s")
	// change the hour marker format from "h" to hourFormat
	if hourFormat != "" {
		tda = strings.Replace(tda, "h", hourFormat, 1)
	}
	// change the minute marker format from "m" to minuteFormat
	if minuteFormat != "" {
		tda = strings.Replace(tda, "m", minuteFormat, 1)
	}
	return tda
}
