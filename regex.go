package tripismapiutilities

import (
	"regexp"
	"strings"
)

// Tokenize simply turns an input string into it's lower case equivalent
func Tokenize(s string) string {
	s = strings.TrimSpace(s)
	return strings.ToLower(s)
}

// TokenizeAndRemoveSpaces turns an input string into it's lower case equivalent and removes any spaces from within the string
func TokenizeAndRemoveSpaces(s string) string {
	if s == "" {
		return s
	}
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	s = strings.Replace(s, " ", "", -1)
	return s
}

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
		s_token := Tokenize(s)
		if strings.HasPrefix(s_token, "a") {
			return true
		}
		if strings.HasPrefix(s_token, "e") {
			return true
		}
		if strings.HasPrefix(s_token, "i") {
			return true
		}
		if strings.HasPrefix(s_token, "o") {
			return true
		}
		if strings.HasPrefix(s_token, "u") {
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

// RemoveNonAlphaNumeric is a regex to remove any non-alphanumeric characters from a string
func RemoveNonAlphaNumeric(s string) string {
	// Make a Regex to say we only want alphanumeric characters
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return s
	}
	return reg.ReplaceAllString(s, "")
}
