package tripismapiutilities

import (
	"strings"
)

// ReplaceMisspellings checks commonly misspelt words and automatically replaces them in search
func ReplaceMisspellings(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	terms := strings.Split(s, " ")

	searchTerm := ""
	for _, term := range terms {
		switch term {
		case "raddiss":
			term = "radiss"
		case "raddisson":
			term = "radisson"
		case "raddison":
			term = "radisson"
		case "radison":
			term = "radisson"
		case "marriot":
			term = "marriott"
		case "mariott":
			term = "marriott"
		case "mariot":
			term = "marriott"
		}
		searchTerm = searchTerm + " " + term
	}
	return searchTerm
}
