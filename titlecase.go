package tripismapiutilities

import "strings"

// TitleCase returns title cased string :) all letters that begin words mapped to their title case.
// Certain small words are skipped to ensure grammatical correctness.
// e.g. WELCOME TO THE DOLLHOUSE => Welcome to the Dollhouse
func TitleCase(s string) string {
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
