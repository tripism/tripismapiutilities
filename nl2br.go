package tripismapiutilities

import "strings"

// Nl2br replaces newlines with HTML breaks (<br/>)
func Nl2br(s string) string {
	return strings.Replace(s, "\n", "<br/>", -1)
}
