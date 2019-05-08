package tripismapiutilities

import "strings"

// TitleCase returns title cased string :) all letters that begin words mapped to their title case.
// Certain small words are skipped to ensure grammatical correctness.
// e.g. WELCOME TO THE DOLLHOUSE => Welcome to the Dollhouse
func TitleCase(s string) string {
	if len(s) == 0 {
		return s
	}
	words := strings.Fields(s)
	smallwords := " a an on the to of in and "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[index] = strings.ToLower(word)
		} else {
			words[index] = strings.Title(strings.ToLower(word))
		}
	}
	return strings.Join(words, " ")
}

// FirstCharacters returns only the specified number of characters starting from the first letter of the provided string.
func FirstCharacters(s string, i int) string {
	if len(s) == 0 {
		return s
	}
	if len(s) < i {
		return s
	}
	return s[:i]
}

// LastCharacters returns only the specified number of characters starting from the last letter of the provided string.
func LastCharacters(s string, i int) string {
	if len(s) == 0 {
		return s
	}
	if len(s) < i {
		return s
	}
	return s[i:]
}
