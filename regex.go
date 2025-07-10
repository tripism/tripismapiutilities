package tripismapiutilities

import (
	"regexp"
	"strings"
)

// RegExSearchAny generates a fuzzy RegEx search pattern for given search text
func RegExSearchAny(search string) string {
	// Split search string on spaces
	ptnS := "(?=.*"
	ptnE := ")"
	ptn := ""
	// Split input string on whitespace
	searchArr := strings.Split(search, " ")
	// Iterate through the input string and create RegEx pattern(s)
	for i := range searchArr {
		if len(searchArr[i]) > 0 {
			ptn += ptnS + searchArr[i] + ptnE
		}
	}
	// Return the compiled RegEx pattern
	return (ptn)
}

// RegExSearchOr generates a fuzzy RegEx search pattern for given search text to search across multiple values
func RegExSearchOr(search string) string {
	// Split search string on spaces
	ptnS := "(?=.*"
	ptnE := ")"
	ptn := ""
	ptnOR := "|"
	// Split input string on whitespace
	searchArr := strings.Split(search, " ")
	// Iterate through the input string and create RegEx pattern(s)
	for i := range searchArr {
		if len(searchArr[i]) > 0 {
			ptn += ptnS + searchArr[i] + ptnE

			if i+1 < len(searchArr) && len(searchArr[i+1]) > 0 {
				ptn += ptnOR
			}
		}
	}
	// Return the compiled RegEx pattern
	return (ptn)
}

// RegExSearchAll generates a fuzzy RegEx search pattern for given search text where all match
func RegExSearchAll(search string) string {
	// Split search string on spaces
	ptnS := "(?=.*"
	ptnE := ")"
	ptn := ptnS + search + ptnE
	// Return the compiled RegEx pattern
	return (ptn)
}

// RegExSearchAllStartsWith generates a fuzzy RegEx search pattern for given search text where start of word matches
func RegExSearchAllStartsWith(search string) string {
	// Split search string on spaces
	ptnS := "(^"
	ptnE := ")"
	ptn := ptnS + search + ptnE
	// Return the compiled RegEx pattern
	return (ptn)
}

// StartsWithVowel checks if a string starts with a voewel
func StartsWithVowel(s string) bool {
	if len(s) > 0 {
		sToken := Tokenize(s)
		if strings.HasPrefix(sToken, "a") {
			return true
		}
		if strings.HasPrefix(sToken, "e") {
			return true
		}
		if strings.HasPrefix(sToken, "i") {
			return true
		}
		if strings.HasPrefix(sToken, "o") {
			return true
		}
		if strings.HasPrefix(sToken, "u") {
			return true
		}
	}
	return false
}

// AdjustForVowel returns a word to append if a given word if the word starts with a vowel
func AdjustForVowel(word string, appendword string) string {
	if len(word) > 0 {
		if StartsWithVowel(word) {
			return appendword
		}
	}
	return ""
}

var nonAlphaNumericRegexp = regexp.MustCompile("[^a-zA-Z0-9]+")

// RemoveNonAlphaNumeric is a regex to remove any non-alphanumeric characters from a string
func RemoveNonAlphaNumeric(s string) string {
	return nonAlphaNumericRegexp.ReplaceAllString(s, "")
}
