package tripismapiutilities

import (
	"strings"
)

// AntiCSVInjection is a helper utility to parse output strings when generateing CSV/Excel files for download
// See https://www.owasp.org/index.php/CSV_Injection
func AntiCSVInjection(cellContents string) string {
	if strings.HasPrefix(cellContents, "=") {
		return "'" + cellContents
	}
	if strings.HasPrefix(cellContents, "+") {
		return "'" + cellContents
	}
	if strings.HasPrefix(cellContents, "-") {
		return "'" + cellContents
	}
	if strings.HasPrefix(cellContents, "@") {
		return "'" + cellContents
	}
	return cellContents
}
